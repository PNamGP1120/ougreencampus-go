package content

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
	var req CreateContentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	authorID := c.GetString("user_id")
	if err := h.service.Create(authorID, req); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Message(c, http.StatusCreated, "content created")
}

func (h *Handler) GetAll(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, data)
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")

	var req UpdateContentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.service.Update(id, req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Message(c, http.StatusOK, "content updated")
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Message(c, http.StatusOK, "content deleted")
}
