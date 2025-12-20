package media

import (
	"net/http"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Upload(c *gin.Context) {
	userID, _ := c.Get("user_id")

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, "file is required")
		return
	}
	defer file.Close()

	media, err := h.service.Upload(file, userID.(uint))
	if err != nil {
		common.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SuccessResponse(c, UploadResponse{
		ID:  media.ID,
		URL: media.URL,
	})
}
