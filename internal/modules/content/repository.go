package content

import "gorm.io/gorm"

type Repository interface {
	Create(c *Content) error
	FindAll() ([]Content, error)
	FindByID(id string) (*Content, error)
	Update(c *Content) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(c *Content) error {
	return r.db.Create(c).Error
}

func (r *repository) FindAll() ([]Content, error) {
	var list []Content
	err := r.db.Order("created_at desc").Find(&list).Error
	return list, err
}

func (r *repository) FindByID(id string) (*Content, error) {
	var c Content
	err := r.db.First(&c, "id = ?", id).Error
	return &c, err
}

func (r *repository) Update(c *Content) error {
	return r.db.Save(c).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&Content{}, "id = ?", id).Error
}
