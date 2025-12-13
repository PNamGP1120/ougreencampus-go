package system

import "gorm.io/gorm"

type SystemConfig struct {
	Key   string `gorm:"primaryKey"`
	Value string `gorm:"type:text"`
}

type ConfigRepository interface {
	Set(cfg *SystemConfig) error
	GetAll() ([]SystemConfig, error)
}

type configRepo struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) ConfigRepository {
	return &configRepo{db: db}
}

func (r *configRepo) Set(cfg *SystemConfig) error {
	return r.db.Save(cfg).Error
}

func (r *configRepo) GetAll() ([]SystemConfig, error) {
	var cfgs []SystemConfig
	err := r.db.Find(&cfgs).Error
	return cfgs, err
}
