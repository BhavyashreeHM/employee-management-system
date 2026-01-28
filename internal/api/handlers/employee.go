package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Context is the most important part of gin. It allows us to pass variables
// between middleware, manage the flow, validate the JSON of a request and render a JSON response for example.
func Employee(c *gin.Context) {
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
