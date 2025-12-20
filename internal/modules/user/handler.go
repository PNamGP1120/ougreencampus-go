package user

import (
	"strconv"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
	cfg     *config.Config
}

func NewHandler(service Service, cfg *config.Config) *Handler {
	return &Handler{service: service, cfg: cfg}
}

// ===== AUTH =====

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ErrorResponse(c, 400, err.Error())
		return
	}
	user, err := h.service.Register(req)
	if err != nil {
		common.ErrorResponse(c, 400, err.Error())
		return
	}
	common.SuccessResponse(c, user)
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ErrorResponse(c, 400, err.Error())
		return
	}
	res, err := h.service.Login(req, h.cfg.JWTSecret)
	if err != nil {
		common.ErrorResponse(c, 401, err.Error())
		return
	}
	common.SuccessResponse(c, res)
}

// ===== USER =====

func (h *Handler) Me(c *gin.Context) {
	id := c.GetUint("user_id")
	user, _ := h.service.GetByID(id)
	common.SuccessResponse(c, user)
}

func (h *Handler) UpdateMe(c *gin.Context) {
	id := c.GetUint("user_id")
	var req UpdateProfileRequest
	c.ShouldBindJSON(&req)
	h.service.UpdateProfile(id, req)
	common.SuccessResponse(c, gin.H{"message": "updated"})
}

func (h *Handler) ChangePassword(c *gin.Context) {
	id := c.GetUint("user_id")
	var req UpdatePasswordRequest
	c.ShouldBindJSON(&req)
	h.service.UpdatePassword(id, req)
	common.SuccessResponse(c, gin.H{"message": "password updated"})
}

func (h *Handler) List(c *gin.Context) {
	users, _ := h.service.List()
	common.SuccessResponse(c, users)
}

func (h *Handler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := h.service.GetByID(uint(id))
	common.SuccessResponse(c, user)
}

func (h *Handler) UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req UpdateRoleRequest
	c.ShouldBindJSON(&req)
	h.service.UpdateRole(uint(id), req.Role)
	common.SuccessResponse(c, gin.H{"message": "role updated"})
}

func (h *Handler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req UpdateStatusRequest
	c.ShouldBindJSON(&req)
	h.service.UpdateStatus(uint(id), req.Status)
	common.SuccessResponse(c, gin.H{"message": "status updated"})
}
