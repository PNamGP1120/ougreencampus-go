package middleware

import (
	"net/http"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/gin-gonic/gin"
)

func RequireRoles(roles ...string) gin.HandlerFunc {
	roleSet := map[string]bool{}
	for _, r := range roles {
		roleSet[r] = true
	}

	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok {
			common.ErrorResponse(c, http.StatusForbidden, "forbidden")
			c.Abort()
			return
		}

		if !roleSet[role.(string)] {
			common.ErrorResponse(c, http.StatusForbidden, "permission denied")
			c.Abort()
			return
		}

		c.Next()
	}
}
