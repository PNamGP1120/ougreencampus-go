package content

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler, authMW gin.HandlerFunc, adminMW gin.HandlerFunc) {
	r.GET("/contents", h.ListContent)
	r.GET("/contents/:id", h.GetContent)

	r.POST("/contents", authMW, h.CreateContent)
	r.PUT("/contents/:id", authMW, h.UpdateContent)
	r.DELETE("/contents/:id", authMW, h.DeleteContent)

	r.GET("/categories", h.ListCategories)
	r.POST("/categories", authMW, adminMW, h.CreateCategory)

	r.GET("/tags", h.ListTags)
	r.POST("/tags", authMW, adminMW, h.CreateTag)

	r.POST("/media/upload", authMW, h.UploadMedia)
	r.GET("/media", authMW, h.ListMedia)
	r.DELETE("/media/:id", authMW, adminMW, h.DeleteMedia)
	r.POST("/media/:id/attach", authMW, h.AttachMedia)
}
