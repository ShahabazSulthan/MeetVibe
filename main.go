package main

import (
	"net/http"
	"video-chat-app/server" // Ensure this matches the actual path of your `server` package

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Allow CORS
	r.Use(cors.Default())

	// Load HTML templates for rendering the UI
	r.LoadHTMLGlob("templates/*")

	// Serve the main page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Serve the room page, passing the room ID dynamically
	r.GET("/room/:roomID", func(c *gin.Context) {
		roomID := c.Param("roomID")
		c.HTML(http.StatusOK, "room.html", gin.H{"RoomID": roomID})
	})

	// Initialize the global RoomMap
	server.AllRooms.Init()

	// Routes for creating and joining rooms
	r.POST("/create", server.CreateRoomRequestHandler)
	r.GET("/join", server.JoinRoomRequestHandler)

	// Start the server
	if err := r.Run(":8000"); err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
