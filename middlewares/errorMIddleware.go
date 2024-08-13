package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware that handles errors globally
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request

		// Check if there were any errors
		if len(c.Errors) > 0 {
			// Log errors or perform any error-related logic here
			for _, e := range c.Errors {
				// Send a response with an appropriate status code and message
				c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
			}
		}
	}
}
