package system

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

/* ========= CONFIG ========= */

func (r *Repository) GetConfigs() ([]SystemConfig, error) {
	var items []SystemConfig
	return items, r.db.Find(&items).Error
}

func (r *Repository) SaveConfig(key, value string) error {
	return r.db.Save(&SystemConfig{
		Key:   key,
		Value: value,
	}).Error
}

/* ========= AUDIT ========= */

func (r *Repository) ListAudit(f AuditFilter) ([]AuditLog, int64) {
	var items []AuditLog
	var total int64

	q := r.db.Model(&AuditLog{})

	if f.UserID != "" {
		q = q.Where("user_id = ?", f.UserID)
	}
	if f.Action != "" {
		q = q.Where("action = ?", f.Action)
	}

	q.Count(&total)

	if f.Page > 0 && f.Limit > 0 {
		q = q.Offset((f.Page - 1) * f.Limit).Limit(f.Limit)
	}

	q.Order("created_at desc").Find(&items)
	return items, total
}

/* ========= NOTIFICATION ========= */

func (r *Repository) ListNotifications(userID string, f NotificationFilter) ([]Notification, int64) {
	var items []Notification
	var total int64

	q := r.db.Model(&Notification{}).Where("user_id = ?", userID)
	q.Count(&total)

	if f.Page > 0 && f.Limit > 0 {
		q = q.Offset((f.Page - 1) * f.Limit).Limit(f.Limit)
	}

	q.Order("created_at desc").Find(&items)
	return items, total
}

func (r *Repository) MarkRead(id string) error {
	return r.db.Model(&Notification{}).
		Where("id = ?", id).
		Update("is_read", true).Error
}
