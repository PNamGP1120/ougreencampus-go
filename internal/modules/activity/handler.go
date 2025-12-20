package activity

import (
	"net/http"
	"strconv"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/gin-gonic/gin"
)

/* ---------- ACTIVITY ---------- */

func (h *Handler) ListActivities(c *gin.Context) {
	items, _ := h.repo.ListActivities()
	c.JSON(http.StatusOK, common.Success(items))
}

func (h *Handler) CreateActivity(c *gin.Context) {
	var req CreateActivityRequest
	c.ShouldBindJSON(&req)
	uid, _ := strconv.Atoi(c.GetString("user_id"))

	a := Activity{
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
		Image:       req.Image,
		OrganizerID: uint(uid),
	}
	h.repo.CreateActivity(&a)
	c.JSON(http.StatusOK, common.Success(gin.H{"id": a.ID}))
}

func (h *Handler) GetActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	a, _ := h.repo.GetActivity(uint(id))
	c.JSON(http.StatusOK, common.Success(a))
}

func (h *Handler) UpdateActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	a, _ := h.repo.GetActivity(uint(id))
	var req UpdateActivityRequest
	c.ShouldBindJSON(&req)
	a.Title = req.Title
	a.Image = req.Image
	h.repo.UpdateActivity(a)
	c.JSON(http.StatusOK, common.Message("updated"))
}

func (h *Handler) DeleteActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.repo.DeleteActivity(uint(id))
	c.JSON(http.StatusOK, common.Message("deleted"))
}

/* ---------- PARTICIPANT ---------- */

func (h *Handler) JoinActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid, _ := strconv.Atoi(c.GetString("user_id"))
	h.repo.JoinActivity(uint(uid), uint(id))
	c.JSON(http.StatusOK, common.Message("joined"))
}

func (h *Handler) LeaveActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid, _ := strconv.Atoi(c.GetString("user_id"))
	h.repo.LeaveActivity(uint(uid), uint(id))
	c.JSON(http.StatusOK, common.Message("left"))
}

func (h *Handler) ListParticipants(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	items, _ := h.repo.ListParticipants(uint(id))
	c.JSON(http.StatusOK, common.Success(items))
}

/* ---------- CONTEST ---------- */

func (h *Handler) Submit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid, _ := strconv.Atoi(c.GetString("user_id"))
	var req SubmitRequest
	c.ShouldBindJSON(&req)

	s := Submission{
		UserID:     uint(uid),
		ActivityID: uint(id),
		Content:    req.Content,
		FileURL:    req.FileURL,
		Status:     "pending",
	}
	h.repo.Submit(&s)
	c.JSON(http.StatusOK, common.Success(gin.H{"id": s.ID, "status": s.Status}))
}

func (h *Handler) ListSubmissions(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	items, _ := h.repo.ListSubmissions(uint(id))
	c.JSON(http.StatusOK, common.Success(items))
}

func (h *Handler) ReviewSubmission(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req ReviewSubmissionRequest
	c.ShouldBindJSON(&req)
	h.repo.ReviewSubmission(uint(id), req.Score, req.Comment)
	c.JSON(http.StatusOK, common.Message("reviewed"))
}

/* ---------- CAMPAIGN ---------- */

func (h *Handler) CreateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req CreateTaskRequest
	c.ShouldBindJSON(&req)
	h.repo.CreateTask(&Task{
		ActivityID: uint(id),
		Title:      req.Title,
		Target:     req.Target,
	})
	c.JSON(http.StatusOK, common.Message("task created"))
}

func (h *Handler) ListTasks(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	items, _ := h.repo.ListTasks(uint(id))
	c.JSON(http.StatusOK, common.Success(items))
}

func (h *Handler) AddProgress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid, _ := strconv.Atoi(c.GetString("user_id"))
	var req ProgressRequest
	c.ShouldBindJSON(&req)
	h.repo.AddProgress(&TaskProgress{
		TaskID: uint(id),
		UserID: uint(uid),
		Value:  req.Value,
	})
	c.JSON(http.StatusOK, common.Message("progress added"))
}

/* ---------- PROGRAM ---------- */

func (h *Handler) AddChild(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req AddChildRequest
	c.ShouldBindJSON(&req)
	h.repo.AddChild(uint(id), req.ChildID)
	c.JSON(http.StatusOK, common.Message("child added"))
}

func (h *Handler) ListChildren(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	items, _ := h.repo.ListChildren(uint(id))
	c.JSON(http.StatusOK, common.Success(items))
}
