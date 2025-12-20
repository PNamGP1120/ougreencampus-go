package common

import "github.com/gin-gonic/gin"

// SuccessResponse returns success JSON
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

// ErrorResponse returns error JSON
func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"error":   message,
	})
}
