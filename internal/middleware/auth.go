package middleware

import (
	"net/http"
	"strings"

	"github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
	"github.com/gin-gonic/gin"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func Auth(jwtSvc *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing token"})
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token format"})
			return
		}

		token, err := jwtSvc.ValidateToken(parts[1])
		if err != nil || token == nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		// ✅ jwt/v5: claims mặc định là jwt.MapClaims
		claims, ok := token.Claims.(jwtv5.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token claims"})
			return
		}

		uid, _ := claims["user_id"].(string)
		role, _ := claims["role"].(string)

		if uid == "" || role == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token payload"})
			return
		}

		c.Set("user_id", uid)
		c.Set("role", role)

		c.Next()
	}
}
