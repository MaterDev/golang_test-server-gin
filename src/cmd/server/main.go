package main

import (
	"github.com/gin-gonic/gin"
	"github.com/materdev/golang_test-server-gin/src/api/routes"
)

func main() {
	r := gin.Default()
	
	// Handlers
	routes.RegisterHealthRoutes(r)

	r.Run() // Listen to default port 8080
}