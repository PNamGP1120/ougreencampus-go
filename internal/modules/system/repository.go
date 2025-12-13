package system

import "gorm.io/gorm"

type Repository interface {
	// Config
	UpsertConfig(c *SystemConfig) error
	ListConfigs() ([]SystemConfig, error)
	DeleteConfig(key string) error

	// Audit
	CreateAudit(a *AuditLog) error
	ListAudit(limit int) ([]AuditLog, error)

	// Reports
	CountTable(table string) (int64, error)
}

type repository struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repository{db: db} }

// Config
func (r *repository) UpsertConfig(c *SystemConfig) error {
	return r.db.Save(c).Error
}
func (r *repository) ListConfigs() ([]SystemConfig, error) {
	var list []SystemConfig
	err := r.db.Order("key asc").Find(&list).Error
	return list, err
}
func (r *repository) DeleteConfig(key string) error {
	return r.db.Delete(&SystemConfig{}, "key = ?", key).Error
}

// Audit
func (r *repository) CreateAudit(a *AuditLog) error { return r.db.Create(a).Error }
func (r *repository) ListAudit(limit int) ([]AuditLog, error) {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	var list []AuditLog
	err := r.db.Order("created_at desc").Limit(limit).Find(&list).Error
	return list, err
}

// Reports
func (r *repository) CountTable(table string) (int64, error) {
	var c int64
	err := r.db.Table(table).Count(&c).Error
	return c, err
}
