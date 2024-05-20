package utils

import "github.com/gin-gonic/gin"

// RespondWithError sends a JSON response with an error message and status code
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

// RespondWithJSON sends a JSON response with the given payload and status code
func RespondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
