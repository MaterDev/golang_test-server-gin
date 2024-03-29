package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/materdev/golang_test-server-gin/src/api/handlers" // This imports the handlers package where your logic is
)

/*
This function registers the routes for the healthcheck endpoints.
*/
func RegisterHealthRoutes(router *gin.Engine) {
	router.GET("/allhealthcheck", handlers.AllHealthCheck)
}