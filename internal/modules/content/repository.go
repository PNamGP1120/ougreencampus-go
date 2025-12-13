package content

import "gorm.io/gorm"

type Repository interface {
	Create(content *Content) error
	FindAll() ([]Content, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(content *Content) error {
	return r.db.Create(content).Error
}

func (r *repository) FindAll() ([]Content, error) {
	var contents []Content
	err := r.db.Order("created_at desc").Find(&contents).Error
	return contents, err
}
