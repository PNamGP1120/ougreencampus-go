package router

import (
	"github.com/gin-gonic/gin"

	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
)

type Handlers struct {
	User     *user.Handler
	Content  *content.Handler
	Activity *activity.Handler
	Event    *event.Handler
	System   *system.Handler
}

func SetupRouter(h *Handlers, jwtSecret string) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	authMW := middleware.Auth(jwtSecret)
	adminMW := middleware.RequireRole("admin")

	user.RegisterRoutes(api, h.User, authMW, adminMW)
	content.RegisterRoutes(api, h.Content, authMW, adminMW)
	activity.RegisterRoutes(api, h.Activity, authMW, adminMW)
	event.RegisterRoutes(api, h.Event, authMW, adminMW)
	system.RegisterRoutes(api, h.System, authMW, adminMW)

	return r
}
