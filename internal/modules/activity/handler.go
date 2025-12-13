package activity

import (
	"net/http"

	"github.com/PNamGP1120/ougreencampus-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Create(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.CreateActivity(req.Name, req.Type, c.GetString("user_id")); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Message(c, http.StatusCreated, "activity created")
}

func (h *Handler) List(c *gin.Context) {
	data, err := h.svc.ListActivities()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, data)
}

func (h *Handler) SubmitContest(c *gin.Context) {
	var req struct {
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.SubmitContest(c.Param("id"), c.GetString("user_id"), req.Content); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Message(c, http.StatusOK, "submitted")
}

func (h *Handler) Submissions(c *gin.Context) {
	data, err := h.svc.ListSubmissions(c.Param("id"))
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, data)
}
