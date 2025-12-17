package event

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler, authMW gin.HandlerFunc, adminMW gin.HandlerFunc) {
	events := r.Group("/events")
	{
		events.GET("", h.ListEvents)
		events.GET("/:id", h.GetEvent)

		events.POST("", authMW, h.CreateEvent)
		events.PUT("/:id", authMW, h.UpdateEvent)
		events.DELETE("/:id", authMW, h.DeleteEvent)

		events.POST("/:id/register", authMW, h.Register)
		events.GET("/:id/registrations", authMW, h.Registrations)

		events.POST("/:id/checkin", authMW, h.CheckinQR)
		events.POST("/:id/checkin/manual", authMW, h.CheckinManual)

		events.GET("/:id/stats", authMW, h.Stats)
		events.GET("/:id/export", authMW, h.Export)
	}
}
