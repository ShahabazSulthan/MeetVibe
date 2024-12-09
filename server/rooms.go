package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Participant describes a single entity in a room
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// RoomMap is the main hashmap [roomID string] -> []Participant
type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

// Init initializes the RoomMap struct
func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

// CreateRoom generates a unique room ID and initializes it in the hashmap
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	const roomIDLength = 8
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	b := make([]rune, roomIDLength)
	for i := range b {
		b[i] = rune(charset[rand.Intn(len(charset))])
	}

	roomID := string(b)
	r.Map[roomID] = []Participant{}

	return roomID
}

// InsertIntoRoom adds a participant to a room
func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{Host: host, Conn: conn}
	r.Map[roomID] = append(r.Map[roomID], p)
	log.Printf("Added participant to room %s. Total participants: %d\n", roomID, len(r.Map[roomID]))
}

// RemoveParticipant removes a specific participant from a room
func (r *RoomMap) RemoveParticipant(roomID string, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	if participants, exists := r.Map[roomID]; exists {
		for i, p := range participants {
			if p.Conn == conn {
				r.Map[roomID] = append(participants[:i], participants[i+1:]...)
				break
			}
		}

		if len(r.Map[roomID]) == 0 {
			delete(r.Map, roomID)
			log.Printf("Room %s deleted as it became empty\n", roomID)
		}
	}
}

// Broadcast sends a message to all participants in a room
func (r *RoomMap) Broadcast(roomID string, message []byte, sender *websocket.Conn) {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	if participants, exists := r.Map[roomID]; exists {
		for _, p := range participants {
			if p.Conn != sender {
				if err := p.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Println("Broadcast error:", err)
				}
			}
		}
	}
}

// StartCleaning runs a periodic task to remove inactive connections
func (r *RoomMap) StartCleaning(interval time.Duration) {
	go func() {
		for range time.Tick(interval) {
			r.cleanInactiveConnections()
		}
	}()
}

func (r *RoomMap) cleanInactiveConnections() {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	for roomID, participants := range r.Map {
		activeParticipants := participants[:0]
		for _, p := range participants {
			if err := p.Conn.WriteMessage(websocket.PingMessage, nil); err == nil {
				activeParticipants = append(activeParticipants, p)
			}
		}
		r.Map[roomID] = activeParticipants
		if len(activeParticipants) == 0 {
			delete(r.Map, roomID)
		}
	}
}
