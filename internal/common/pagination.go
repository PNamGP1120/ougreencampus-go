package common

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func NewPagination(page, limit, total int) Pagination {
	return Pagination{
		Page:  page,
		Limit: limit,
		Total: total,
	}
}
