package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

// GetPagination extracts page & limit from query
func GetPagination(c *gin.Context) (page, limit int) {
	page = 1
	limit = 10

	if p := c.Query("page"); p != "" {
		fmt.Sscan(p, &page)
	}
	if l := c.Query("limit"); l != "" {
		fmt.Sscan(l, &limit)
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return
}
