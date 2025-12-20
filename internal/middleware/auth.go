package middleware

import (
	"net/http"
	"strings"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	jwtutil "github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			common.ErrorResponse(c, http.StatusUnauthorized, "missing token")
			c.Abort()
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			common.ErrorResponse(c, http.StatusUnauthorized, "invalid token format")
			c.Abort()
			return
		}

		claims, err := jwtutil.ParseToken(parts[1], cfg.JWTSecret)
		if err != nil {
			common.ErrorResponse(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
