package activity

import (
	"strconv"
	"time"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{svc: s}
}

// 30 GET /activities
func (h *Handler) List(c *gin.Context) {
	page, limit := common.GetPagination(c)

	search := c.Query("search")
	typ := c.Query("type")
	status := c.Query("status")

	var from, to *time.Time
	if v := c.Query("from"); v != "" {
		t, _ := time.Parse(time.RFC3339, v)
		from = &t
	}
	if v := c.Query("to"); v != "" {
		t, _ := time.Parse(time.RFC3339, v)
		to = &t
	}

	items, total, _ := h.svc.List(search, typ, status, from, to, page, limit)

	common.SuccessResponse(c, gin.H{
		"items": items,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// 31 POST /activities
func (h *Handler) Create(c *gin.Context) {
	var req CreateActivityRequest
	c.ShouldBindJSON(&req)
	a, _ := h.svc.Create(c.GetUint("user_id"), req)
	common.SuccessResponse(c, gin.H{"id": a.ID})
}

// 32 GET /activities/:id
func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	a, _ := h.svc.Get(uint(id))
	common.SuccessResponse(c, a)
}

// 33 PUT /activities/:id
func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req UpdateActivityRequest
	c.ShouldBindJSON(&req)
	h.svc.Update(uint(id), req)
	common.SuccessResponse(c, gin.H{"message": "updated"})
}

// 34 DELETE /activities/:id
func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.Delete(uint(id))
	common.SuccessResponse(c, gin.H{"message": "deleted"})
}

// 35 POST /activities/:id/join
func (h *Handler) Join(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.Join(uint(id), c.GetUint("user_id"))
	common.SuccessResponse(c, gin.H{"message": "joined"})
}

// 36 DELETE /activities/:id/leave
func (h *Handler) Leave(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.Leave(uint(id), c.GetUint("user_id"))
	common.SuccessResponse(c, gin.H{"message": "left"})
}

// 37 GET /activities/:id/participants
func (h *Handler) Participants(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, limit := common.GetPagination(c)
	items, total, _ := h.svc.Participants(uint(id), page, limit)

	common.SuccessResponse(c, gin.H{
		"items": items,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// 38 POST /activities/:id/submissions
func (h *Handler) Submit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req SubmitContestRequest
	c.ShouldBindJSON(&req)
	sub, _ := h.svc.Submit(uint(id), c.GetUint("user_id"), req)
	common.SuccessResponse(c, gin.H{"id": sub.ID, "status": sub.Status})
}

// 39 GET /activities/:id/submissions
func (h *Handler) Submissions(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, limit := common.GetPagination(c)
	status := c.Query("status")

	items, total, _ := h.svc.Submissions(uint(id), status, page, limit)
	common.SuccessResponse(c, gin.H{
		"items": items,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// 40 PATCH /submissions/:id/review
func (h *Handler) Review(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req ReviewSubmissionRequest
	c.ShouldBindJSON(&req)
	h.svc.Review(uint(id), req.Score, req.Comment)
	common.SuccessResponse(c, gin.H{"message": "reviewed"})
}

// 41 POST /activities/:id/tasks
func (h *Handler) CreateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req CreateTaskRequest
	c.ShouldBindJSON(&req)
	task, _ := h.svc.CreateTask(uint(id), req)
	common.SuccessResponse(c, gin.H{"id": task.ID})
}

// 42 GET /activities/:id/tasks
func (h *Handler) Tasks(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, limit := common.GetPagination(c)
	items, total, _ := h.svc.Tasks(uint(id), page, limit)

	common.SuccessResponse(c, gin.H{
		"items": items,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// 43 POST /tasks/:id/progress
func (h *Handler) Progress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req ProgressRequest
	c.ShouldBindJSON(&req)
	h.svc.Progress(uint(id), c.GetUint("user_id"), req.Value)
	common.SuccessResponse(c, gin.H{"message": "progress updated"})
}

// 44 GET /activities/:id/metrics
func (h *Handler) Metrics(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, _ := h.svc.Metrics(uint(id))
	common.SuccessResponse(c, data)
}

// 45 POST /activities/:id/children
func (h *Handler) AddChild(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req AddChildRequest
	c.ShouldBindJSON(&req)
	h.svc.AddChild(uint(id), req.ChildID)
	common.SuccessResponse(c, gin.H{"message": "attached"})
}

// 46 GET /activities/:id/children
func (h *Handler) Children(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	items, _ := h.svc.Children(uint(id))
	common.SuccessResponse(c, items)
}
