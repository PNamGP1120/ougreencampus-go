package middleware

import (
	"strings"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	jwtutil "github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "missing authorization"})
			return
		}

		parts := strings.Split(h, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"message": "invalid authorization"})
			return
		}

		claims, err := jwtutil.ParseToken(parts[1], config.Cfg.JWT.Secret)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "invalid token"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// ===== Helpers dùng cho handler =====

func MustUserID(c *gin.Context) (uint, bool) {
	v, ok := c.Get("user_id")
	if !ok {
		return 0, false
	}
	id, ok := v.(uint)
	return id, ok
}

func MustRole(c *gin.Context) (string, bool) {
	v, ok := c.Get("role")
	if !ok {
		return "", false
	}
	role, ok := v.(string)
	return role, ok
}
