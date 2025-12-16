package user

import (
	"net/http"

	"github.com/PNamGP1120/ougreencampus-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	created, err := h.service.CreateUser(req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, created)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, users)
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")
	u, err := h.service.GetByID(id)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, u)
}

func (h *Handler) UpdateRole(c *gin.Context) {
	id := c.Param("id")

	var req UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.service.UpdateRole(id, req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Message(c, http.StatusOK, "role updated")
}

func (h *Handler) ToggleActive(c *gin.Context) {
	id := c.Param("id")

	var req ToggleActiveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.service.SetActive(id, req.IsActive); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Message(c, http.StatusOK, "user status updated")
}

func (h *Handler) UpdateAvatar(c *gin.Context) {
	id := c.Param("id")

	var req UpdateAvatarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.service.UpdateAvatar(id, req.Avatar); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Message(c, http.StatusOK, "avatar updated")
}
