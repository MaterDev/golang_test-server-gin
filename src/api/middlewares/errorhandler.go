package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware for handling errors encountered during HTTP requests.
func ErrorHandler(c *gin.Context) {
	// This function will be executed before every request to the server
	c.Next()

	// Check if there were any errors during processing
	if len(c.Errors) > 0 {
		// Log the error
		for _, e := range c.Errors {
			Logger.Error(e.Error())
		}

		// * You can add more complex error handling here, such as differentiating between types of errors, and responding differently based on error content or type.

		// Respond with a generic error message and a 500 status code
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "An unexpected error occurred.",
		})
		// Prevent any further response from being sent
		c.Abort()
	}
}