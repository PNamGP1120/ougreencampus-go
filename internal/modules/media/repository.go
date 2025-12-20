package media

import "gorm.io/gorm"

type Repository interface {
	Create(media *Media) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(media *Media) error {
	return r.db.Create(media).Error
}
