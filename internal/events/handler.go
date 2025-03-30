package events

import (
	"net/http"

	"math/rand/v2"

	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.RouterGroup) {
	// Define the route for handling events
	server.GET("/", GetEvent)
	server.POST("/", CreateEvent)
}

func CreateEvent(c *gin.Context) {
	var req Event
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.DateTime = req.DateTime.UTC()
	req.ID = rand.Int()
	req.UserID = rand.Int()
	c.JSON(200, gin.H{
		"message": "Event created successfully",
		"event":   req,
	})
}

// EventHandler handles events for the application.
func GetEvent(c *gin.Context) {
	// Your event handling logic goes here

	c.JSON(200, gin.H{
		"message": "Event handled successfully",
	})
}
