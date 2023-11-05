package main

import (
	"github.com/gin-gonic/gin"
	"github.com/materdev/golang_test-server-gin/src/api/middlewares"
	"github.com/materdev/golang_test-server-gin/src/api/routes"
)

func main() {
	r := gin.Default()

	// Register global error handling middleware
	r.Use(middlewares.ErrorHandler)
	
	// Handlers
	routes.RegisterHealthRoutes(r)

	r.Run() // Listen to default port 8080
}