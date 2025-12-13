package user

import (
	"net/http"

	"github.com/PNamGP1120/ougreencampus-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authSvc AuthService
}

func NewAuthHandler(authSvc AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	token, err := h.authSvc.Login(req.Email, req.Password)
	if err != nil {
		response.Unauthorized(c)
		return
	}

	response.Success(c, http.StatusOK, gin.H{
		"token": token,
	})
}
