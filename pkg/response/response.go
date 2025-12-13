package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, APIResponse{
		Success: true,
		Data:    data,
	})
}

func Message(c *gin.Context, status int, message string) {
	c.JSON(status, APIResponse{
		Success: true,
		Message: message,
	})
}

func Error(c *gin.Context, status int, err interface{}) {
	c.JSON(status, APIResponse{
		Success: false,
		Error:   err,
	})
}

func BadRequest(c *gin.Context, err interface{}) {
	Error(c, http.StatusBadRequest, err)
}

func Unauthorized(c *gin.Context) {
	Error(c, http.StatusUnauthorized, "unauthorized")
}

func Forbidden(c *gin.Context) {
	Error(c, http.StatusForbidden, "forbidden")
}

func InternalServerError(c *gin.Context, err interface{}) {
	Error(c, http.StatusInternalServerError, err)
}
