package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var AllRooms RoomMap

func init() {
	AllRooms.Init()
	go AllRooms.StartCleaning(10 * time.Minute)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// CreateRoomRequestHandler handles the creation of a new room
func CreateRoomRequestHandler(c *gin.Context) {
	roomID := AllRooms.CreateRoom()
	c.JSON(http.StatusOK, gin.H{"room_id": roomID})
}

// JoinRoomRequestHandler handles clients joining a room
func JoinRoomRequestHandler(c *gin.Context) {
	roomID := c.Query("roomID")
	if roomID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomID is required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	AllRooms.InsertIntoRoom(roomID, false, conn)
	defer AllRooms.RemoveParticipant(roomID, conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		AllRooms.Broadcast(roomID, message, conn)
	}
	conn.Close()
}
