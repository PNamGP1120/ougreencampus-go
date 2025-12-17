package user

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

/* ===== AUTH ===== */

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"message": "invalid"})
		return
	}
	u, _ := h.svc.Register(req)
	c.JSON(201, u)
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"message": "invalid"})
		return
	}
	token, u, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token, "user": u})
}

/* ===== USER ===== */

func (h *Handler) ListUsers(c *gin.Context) {
	users, _ := h.svc.ListUsers()
	c.JSON(200, users)
}

func (h *Handler) GetUser(c *gin.Context) {
	u, _ := h.svc.GetUser(c.Param("id"))
	c.JSON(200, u)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	c.ShouldBindJSON(&req)
	u, _ := h.svc.CreateUser(req)
	c.JSON(201, u)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	c.ShouldBindJSON(&req)
	h.svc.UpdateUser(c.Param("id"), req)
	c.JSON(200, gin.H{"message": "updated"})
}

func (h *Handler) ChangeRole(c *gin.Context) {
	var req ChangeRoleRequest
	c.ShouldBindJSON(&req)
	h.svc.ChangeRole(c.Param("id"), req.Role)
	c.JSON(200, gin.H{"message": "role updated"})
}

func (h *Handler) ChangeStatus(c *gin.Context) {
	var req ChangeStatusRequest
	c.ShouldBindJSON(&req)
	h.svc.ChangeStatus(c.Param("id"), req.Status)
	c.JSON(200, gin.H{"message": "status updated"})
}

func (h *Handler) ChangePassword(c *gin.Context) {
	userID := c.GetString("user_id")
	var req ChangePasswordRequest
	c.ShouldBindJSON(&req)
	h.svc.ChangePassword(userID, req.OldPassword, req.NewPassword)
	c.JSON(200, gin.H{"message": "password changed"})
}
