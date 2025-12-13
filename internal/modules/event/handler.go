package event

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

func (h *Handler) Create(c *gin.Context) {
	var req CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	creatorID := c.GetString("user_id")
	if err := h.service.Create(creatorID, req); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Message(c, http.StatusCreated, "event created")
}

func (h *Handler) GetAll(c *gin.Context) {
	list, err := h.service.GetAll()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, list)
}

func (h *Handler) Register(c *gin.Context) {
	eventID := c.Param("id")
	userID := c.GetString("user_id")

	if err := h.service.Register(eventID, userID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Message(c, http.StatusOK, "registered successfully")
}

func (h *Handler) CheckIn(c *gin.Context) {
	eventID := c.Param("id")
	userID := c.GetString("user_id")

	if err := h.service.CheckIn(eventID, userID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Message(c, http.StatusOK, "checked in")
}
