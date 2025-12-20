package middleware

import "github.com/gin-gonic/gin"

func RequireRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := MustRole(c)
		if !ok {
			c.AbortWithStatusJSON(403, gin.H{"message": "forbidden"})
			return
		}

		for _, r := range roles {
			if role == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(403, gin.H{"message": "permission denied"})
	}
}
