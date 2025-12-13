package activity

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
	var req CreateActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	creatorID := c.GetString("user_id")
	if err := h.service.Create(creatorID, req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Message(c, http.StatusCreated, "activity created")
}

func (h *Handler) GetAll(c *gin.Context) {
	acts, err := h.service.GetAll()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, acts)
}

func (h *Handler) Submit(c *gin.Context) {
	activityID := c.Param("id")
	userID := c.GetString("user_id")

	var req SubmitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.service.Submit(activityID, userID, req); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Message(c, http.StatusOK, "submission sent")
}
