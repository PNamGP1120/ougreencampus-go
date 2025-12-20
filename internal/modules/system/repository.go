package system

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

/* ---------- CONFIG ---------- */

func (r *Repository) GetConfigs() ([]SystemConfig, error) {
	var items []SystemConfig
	return items, r.db.Find(&items).Error
}

func (r *Repository) UpsertConfig(key, value string) error {
	var cfg SystemConfig
	err := r.db.Where("key = ?", key).First(&cfg).Error
	if err != nil {
		return r.db.Create(&SystemConfig{Key: key, Value: value}).Error
	}
	cfg.Value = value
	return r.db.Save(&cfg).Error
}

/* ---------- AUDIT ---------- */

func (r *Repository) ListAuditLogs() ([]AuditLog, error) {
	var items []AuditLog
	return items, r.db.Order("created_at desc").Find(&items).Error
}

/* ---------- NOTIFICATION ---------- */

func (r *Repository) ListNotifications(userID uint) ([]Notification, error) {
	var items []Notification
	return items, r.db.Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&items).Error
}

func (r *Repository) MarkRead(id uint) error {
	return r.db.Model(&Notification{}).
		Where("id = ?", id).
		Update("read", true).Error
}
