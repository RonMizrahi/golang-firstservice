package routes

import (
	"restapi/internal/handlers"

	"github.com/gin-gonic/gin"
)

// EventRoutes sets up the routes for event handling.

func EventRoutes(server *gin.RouterGroup) {
	// Define the route for handling events
	server.GET("/event", handlers.GetEvent)
}
