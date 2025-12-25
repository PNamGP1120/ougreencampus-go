package router

import (
	"time"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/media"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.New()

	// 1️⃣ CORS PHẢI ĐỨNG ĐẦU
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 2️⃣ GLOBAL OPTIONS – CHO CORS CHẠY TRƯỚC
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Next()      // 🔴 CỰC KỲ QUAN TRỌNG
		c.Status(204) // trả sau khi CORS gắn header
	})

	// 3️⃣ middleware khác
	r.Use(
		gin.Recovery(),
		middleware.Logger(),
	)

	// 4️⃣ routes
	api := r.Group("/api")
	user.RegisterRoutes(api, db, cfg)
	media.RegisterRoutes(api, db, cfg)
	content.RegisterRoutes(api, db, cfg)
	activity.RegisterRoutes(api, db, cfg)
	event.RegisterRoutes(api, db, cfg)
	system.RegisterRoutes(api, db, cfg)

	return r
}
