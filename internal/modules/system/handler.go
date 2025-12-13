package system

import (
	"net/http"
	"strconv"

	"github.com/PNamGP1120/ougreencampus-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct{ svc Service }

func NewHandler(svc Service) *Handler { return &Handler{svc: svc} }

// Config (Admin)
func (h *Handler) ListConfigs(c *gin.Context) {
	data, err := h.svc.ListConfigs()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, data)
}

func (h *Handler) UpsertConfig(c *gin.Context) {
	var req UpsertConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.svc.UpsertConfig(c.GetString("user_id"), req); err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Message(c, http.StatusOK, "config updated")
}

func (h *Handler) DeleteConfig(c *gin.Context) {
	if err := h.svc.DeleteConfig(c.Param("key")); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Message(c, http.StatusOK, "config deleted")
}

// Reports (Admin)
func (h *Handler) Overview(c *gin.Context) {
	data, err := h.svc.Overview()
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, data)
}

// Audit (Admin)
func (h *Handler) ListAudit(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	data, err := h.svc.ListAudit(limit)
	if err != nil {
		response.InternalServerError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, data)
}
