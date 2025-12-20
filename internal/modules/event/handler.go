package event

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/gin-gonic/gin"
)

/* ---------- EVENT ---------- */

func (h *Handler) ListEvents(c *gin.Context) {
	items, _ := h.repo.ListEvents()
	c.JSON(http.StatusOK, common.Success(items))
}

func (h *Handler) CreateEvent(c *gin.Context) {
	var req CreateEventRequest
	c.ShouldBindJSON(&req)
	uid, _ := strconv.Atoi(c.GetString("user_id"))

	e := Event{
		Title:       req.Title,
		Time:        req.Time,
		Location:    req.Location,
		Image:       req.Image,
		OrganizerID: uint(uid),
	}
	h.repo.CreateEvent(&e)
	c.JSON(http.StatusOK, common.Success(gin.H{"id": e.ID}))
}

func (h *Handler) GetEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	e, _ := h.repo.GetEvent(uint(id))
	c.JSON(http.StatusOK, common.Success(e))
}

func (h *Handler) UpdateEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	e, _ := h.repo.GetEvent(uint(id))
	var req UpdateEventRequest
	c.ShouldBindJSON(&req)
	e.Title = req.Title
	e.Image = req.Image
	h.repo.UpdateEvent(e)
	c.JSON(http.StatusOK, common.Message("updated"))
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.repo.DeleteEvent(uint(id))
	c.JSON(http.StatusOK, common.Message("deleted"))
}

/* ---------- REGISTRATION ---------- */

func (h *Handler) Register(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid, _ := strconv.Atoi(c.GetString("user_id"))
	h.repo.Register(uint(id), uint(uid))
	c.JSON(http.StatusOK, common.Message("registered"))
}

func (h *Handler) ListRegistrations(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	items, _ := h.repo.ListRegistrations(uint(id))
	c.JSON(http.StatusOK, common.Success(items))
}

func (h *Handler) SendConfirmation(c *gin.Context) {
	c.JSON(http.StatusOK, common.Message("confirmation sent"))
}

func (h *Handler) CheckinQR(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid, _ := strconv.Atoi(c.GetString("user_id"))
	h.repo.Checkin(uint(id), uint(uid))
	c.JSON(http.StatusOK, common.Message("checked in"))
}

func (h *Handler) CheckinManual(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req ManualCheckinRequest
	c.ShouldBindJSON(&req)
	h.repo.Checkin(uint(id), req.UserID)
	c.JSON(http.StatusOK, common.Message("checked in"))
}

func (h *Handler) Stats(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	stats, _ := h.repo.Stats(uint(id))
	c.JSON(http.StatusOK, common.Success(stats))
}

func (h *Handler) Export(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	items, _ := h.repo.ListRegistrations(uint(id))

	c.Header("Content-Disposition", "attachment; filename=event.csv")
	c.Header("Content-Type", "text/csv")
	writer := csv.NewWriter(c.Writer)
	writer.Write([]string{"UserID", "Checked"})
	for _, r := range items {
		writer.Write([]string{
			strconv.Itoa(int(r.UserID)),
			strconv.FormatBool(r.Checked),
		})
	}
	writer.Flush()
}
