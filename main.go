package main

import (
	"restapi/internal/events"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	api := server.Group("/api/event") // base path like /api/
	events.EventRoutes(api)

	server.Run("127.0.0.1:8080")
}
