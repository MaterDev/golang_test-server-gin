package main

import (
	"github.com/gin-gonic/gin"
	"github.com/materdev/golang_test-server-gin/src/handlers"
)

func main() {
	r := gin.Default()
	
	// Handlers
	r.GET("/healthcheck", handlers.HealthCheck)

	r.Run() // Listen to default port 8080
}