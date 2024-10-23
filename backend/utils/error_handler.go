package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there were any errors during the request handling
		if len(c.Errors) > 0 {
			// Handle the errors
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": c.Errors,
			})
		}
	}
}
