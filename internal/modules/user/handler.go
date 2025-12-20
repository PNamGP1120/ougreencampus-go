package user

import (
	"net/http"
	"strconv"

	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/PNamGP1120/ougreencampus-go/pkg/hash"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc  *Service
	repo *Repository
}

func NewHandler(s *Service, r *Repository) *Handler {
	return &Handler{svc: s, repo: r}
}

/* ===================== AUTH (1–7) ===================== */

// 1 POST /auth/register
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid input"})
		return
	}
	if err := h.svc.Register(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "registered"})
}

// 2 POST /auth/login
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid input"})
		return
	}

	token, u, err := h.svc.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"token": token,
			"user":  u,
		},
	})
}

// 3 POST /auth/logout
func (h *Handler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "logged out"})
}

// 4 GET /auth/me
func (h *Handler) Me(c *gin.Context) {
	uid, ok := middleware.MustUserID(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "unauthorized"})
		return
	}

	u, err := h.repo.FindByID(uid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "user not found"})
		return
	}

	u.Password = ""
	c.JSON(http.StatusOK, gin.H{"success": true, "data": u})
}

// 5 POST /auth/refresh
func (h *Handler) Refresh(c *gin.Context) {
	uid, ok := middleware.MustUserID(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "unauthorized"})
		return
	}
	role, _ := middleware.MustRole(c)

	token, err := h.svc.RefreshToken(uid, role)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "refresh failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"token": token}})
}

// 6 POST /auth/forgot-password
func (h *Handler) ForgotPassword(c *gin.Context) {
	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid input"})
		return
	}
	h.svc.ForgotPassword(req.Email)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "if email exists, reset link sent"})
}

// 7 POST /auth/reset-password
func (h *Handler) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid input"})
		return
	}
	if err := h.svc.ResetPassword(req.Token, req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "password reset success"})
}

/* ===================== ADMIN (8–13) ===================== */

// 8 GET /users
func (h *Handler) ListUsers(c *gin.Context) {
	users, err := h.repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "cannot list users"})
		return
	}
	for i := range users {
		users[i].Password = ""
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": users})
}

// 9 POST /users
func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid input"})
		return
	}
	if err := h.svc.CreateUser(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "user created"})
}

// 10 GET /users/:id
func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid id"})
		return
	}
	u, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "user not found"})
		return
	}
	u.Password = ""
	c.JSON(http.StatusOK, gin.H{"success": true, "data": u})
}

// 11 PUT /users/:id
func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid id"})
		return
	}
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid input"})
		return
	}

	u, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "user not found"})
		return
	}

	u.Name = req.Name
	u.Avatar = req.Avatar
	_ = h.repo.Update(u)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "updated"})
}

// 12 PATCH /users/:id/role
func (h *Handler) ChangeRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid id"})
		return
	}
	var req ChangeRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid input"})
		return
	}
	u, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "user not found"})
		return
	}
	u.Role = req.Role
	_ = h.repo.Update(u)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "role updated"})
}

// 13 PATCH /users/:id/status
func (h *Handler) ChangeStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid id"})
		return
	}
	var req ChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid input"})
		return
	}
	u, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "user not found"})
		return
	}
	u.Status = req.Status
	_ = h.repo.Update(u)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "status updated"})
}

/* ===================== SELF (14–16) ===================== */

// 14 GET /users/me
func (h *Handler) GetMe(c *gin.Context) {
	h.Me(c)
}

// 15 PUT /users/me
func (h *Handler) UpdateMe(c *gin.Context) {
	uid, ok := middleware.MustUserID(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "unauthorized"})
		return
	}

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid input"})
		return
	}

	u, err := h.repo.FindByID(uid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "user not found"})
		return
	}

	u.Name = req.Name
	u.Avatar = req.Avatar
	_ = h.repo.Update(u)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "profile updated"})
}

// 16 PUT /users/me/password
func (h *Handler) ChangePassword(c *gin.Context) {
	uid, ok := middleware.MustUserID(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "unauthorized"})
		return
	}

	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid input"})
		return
	}

	u, err := h.repo.FindByID(uid)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "user not found"})
		return
	}

	if !hash.CheckPassword(u.Password, req.OldPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "wrong old password"})
		return
	}

	newHash, _ := hash.HashPassword(req.NewPassword)
	u.Password = newHash
	_ = h.repo.Update(u)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "password changed"})
}
