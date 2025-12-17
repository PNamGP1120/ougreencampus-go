package activity

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler, authMW gin.HandlerFunc, adminMW gin.HandlerFunc) {
	activities := r.Group("/activities")
	{
		activities.GET("", h.ListActivities)
		activities.GET("/:id", h.GetActivity)

		activities.POST("", authMW, h.CreateActivity)
		activities.PUT("/:id", authMW, h.UpdateActivity)
		activities.DELETE("/:id", authMW, h.DeleteActivity)

		activities.POST("/:id/join", authMW, h.Join)
		activities.DELETE("/:id/leave", authMW, h.Leave)

		activities.GET("/:id/participants", authMW, h.Participants)

		activities.POST("/:id/submissions", authMW, h.Submit)
		activities.GET("/:id/submissions", authMW, h.Submissions)

		activities.POST("/:id/tasks", authMW, h.CreateTask)
		activities.GET("/:id/tasks", authMW, h.Tasks)

		activities.POST("/:id/children", authMW, h.AddChild)
		activities.GET("/:id/children", h.Children)
	}

	r.PATCH("/submissions/:id/review", authMW, adminMW, h.ReviewSubmission)
	r.POST("/tasks/:id/progress", authMW, h.Progress)
}
