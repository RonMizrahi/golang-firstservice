package main

import (
	"restapi/internal/events"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	eventGroup := server.Group("/api/event") // base path like /api/
	events.EventRoutes(eventGroup)

	server.Run("127.0.0.1:8080")
}
