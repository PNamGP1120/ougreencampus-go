package activity

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{svc: s}
}

/* ========= ACTIVITY ========= */

func (h *Handler) ListActivities(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	filter := ListActivityFilter{
		Search: c.Query("search"),
		Type:   c.Query("type"),
		Status: c.Query("status"),
		Page:   page,
		Limit:  limit,
	}

	items, total := h.svc.ListActivities(filter)

	c.JSON(200, gin.H{
		"items": items,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) GetActivity(c *gin.Context) {
	data, _ := h.svc.GetActivity(c.Param("id"))
	c.JSON(200, data)
}

func (h *Handler) CreateActivity(c *gin.Context) {
	var req CreateActivityRequest
	_ = c.ShouldBindJSON(&req)
	id, _ := h.svc.CreateActivity(req, c.GetString("user_id"))
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) UpdateActivity(c *gin.Context) {
	var req UpdateActivityRequest
	_ = c.ShouldBindJSON(&req)
	_ = h.svc.UpdateActivity(c.Param("id"), req)
	c.JSON(200, gin.H{"message": "updated"})
}

func (h *Handler) DeleteActivity(c *gin.Context) {
	_ = h.svc.DeleteActivity(c.Param("id"))
	c.JSON(200, gin.H{"message": "deleted"})
}

/* ========= PARTICIPANT ========= */

func (h *Handler) Join(c *gin.Context) {
	_ = h.svc.Join(c.Param("id"), c.GetString("user_id"))
	c.JSON(200, gin.H{"message": "joined"})
}

func (h *Handler) Leave(c *gin.Context) {
	_ = h.svc.Leave(c.Param("id"), c.GetString("user_id"))
	c.JSON(200, gin.H{"message": "left"})
}

func (h *Handler) Participants(c *gin.Context) {
	data, _ := h.svc.Participants(c.Param("id"))
	c.JSON(200, data)
}

/* ========= SUBMISSION ========= */

func (h *Handler) Submit(c *gin.Context) {
	var req SubmitRequest
	_ = c.ShouldBindJSON(&req)
	id, _ := h.svc.Submit(c.Param("id"), c.GetString("user_id"), req)
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) Submissions(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	data, total := h.svc.Submissions(c.Param("id"), c.Query("status"), page, limit)
	c.JSON(200, gin.H{"items": data, "total": total})
}

func (h *Handler) ReviewSubmission(c *gin.Context) {
	var req ReviewSubmissionRequest
	_ = c.ShouldBindJSON(&req)
	_ = h.svc.ReviewSubmission(c.Param("id"), req)
	c.JSON(200, gin.H{"message": "reviewed"})
}

/* ========= CAMPAIGN ========= */

func (h *Handler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest
	_ = c.ShouldBindJSON(&req)
	id, _ := h.svc.CreateTask(c.Param("id"), req)
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) Tasks(c *gin.Context) {
	data, _ := h.svc.Tasks(c.Param("id"))
	c.JSON(200, data)
}

func (h *Handler) Progress(c *gin.Context) {
	var req ProgressRequest
	_ = c.ShouldBindJSON(&req)
	_ = h.svc.Progress(c.Param("id"), c.GetString("user_id"), req.Value)
	c.JSON(200, gin.H{"message": "updated"})
}

/* ========= PROGRAM ========= */

func (h *Handler) AddChild(c *gin.Context) {
	var body struct {
		ChildID string `json:"child_id"`
	}
	_ = c.ShouldBindJSON(&body)
	_ = h.svc.AddChild(c.Param("id"), body.ChildID)
	c.JSON(200, gin.H{"message": "attached"})
}

func (h *Handler) Children(c *gin.Context) {
	data, _ := h.svc.Children(c.Param("id"))
	c.JSON(200, data)
}
