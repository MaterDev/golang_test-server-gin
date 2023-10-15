package main

import (
	"github.com/gin-gonic/gin"
	"github.com/materdev/golang_test-server-gin/src/api/routes"
)

func main() {
	r := gin.Default()
	
	// Handlers
	r.GET("/healthcheck", routes.HealthCheck)

	r.Run() // Listen to default port 8080
}