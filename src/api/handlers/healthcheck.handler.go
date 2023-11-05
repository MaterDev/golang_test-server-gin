package handlers

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/materdev/golang_test-server-gin/src/api/middlewares"

)

/*
This is the handler for the /allhealthcheck route. It's purpose is to check the health of the server and all of its dependencies.

@api {get} /allhealthcheck AllHealthCheck
*/
func AllHealthCheck(c *gin.Context) {
	// Let's force an error when a specific query parameter is set
	if c.Query("forceError") == "true" {
		err := errors.New("forced error for testing")
		c.Error(err) // Add the error to Gin's context
		return       // Return early, the middleware will handle the response
	}

	// Logger message: Healthy
	middlewares.Logger.Info("AllHealthCheck called: server is healthy")
	// JSON Response
	c.JSON(http.StatusOK, gin.H{
		"message": "🙌🏾 OK 🙌🏾",
	})
}