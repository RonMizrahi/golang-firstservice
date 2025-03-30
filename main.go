package main

import (
	"restapi/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	api := server.Group("/api") // base path like /api/event
	routes.EventRoutes(api)

	server.Run("127.0.0.1:8080")
}
