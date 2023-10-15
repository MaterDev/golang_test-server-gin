package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/materdev/golang_test-server-gin/src/api/middlewares"

)

func HealthCheck(c *gin.Context) {
	// Testing the logger
	logger.Info.Println("This is an informational message.")

	// JSON Response
	c.JSON(http.StatusOK, gin.H{
		"message": "🙌🏾 OK 🙌🏾",
	})
}