package system

/* ========= SYSTEM CONFIG ========= */

type UpdateConfigRequest struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

/* ========= AUDIT FILTER ========= */

type AuditFilter struct {
	UserID string
	Action string
	Page   int
	Limit  int
}

/* ========= NOTIFICATION FILTER ========= */

type NotificationFilter struct {
	Page  int
	Limit int
}
