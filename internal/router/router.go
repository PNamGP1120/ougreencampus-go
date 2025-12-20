package router

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/media"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(
		gin.Recovery(),
		middleware.Logger(),
	)

	api := r.Group("/api")

	// Register modules
	user.RegisterRoutes(api, db, cfg)
	media.RegisterRoutes(api, db, cfg)

	return r
}
