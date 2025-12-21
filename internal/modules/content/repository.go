package content

import "gorm.io/gorm"

type Repository interface {
	List(filter map[string]interface{}, page, limit int) ([]Content, int64, error)
	Create(c *Content) error
	FindByID(id uint) (*Content, error)
	Update(c *Content) error
	Delete(id uint) error

	ListCategories() ([]Category, error)
	CreateCategory(c *Category) error

	ListTags() ([]Tag, error)
	CreateTag(t *Tag) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) List(filter map[string]interface{}, page, limit int) ([]Content, int64, error) {
	var items []Content
	var total int64

	q := r.db.Model(&Content{}).Preload("Tags").Preload("Category")

	for k, v := range filter {
		q = q.Where(k, v)
	}

	q.Count(&total)
	err := q.Offset((page - 1) * limit).Limit(limit).Find(&items).Error
	return items, total, err
}

func (r *repository) Create(c *Content) error {
	return r.db.Create(c).Error
}

func (r *repository) FindByID(id uint) (*Content, error) {
	var c Content
	err := r.db.Preload("Tags").Preload("Category").First(&c, id).Error
	return &c, err
}

func (r *repository) Update(c *Content) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(c).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Content{}, id).Error
}

func (r *repository) ListCategories() ([]Category, error) {
	var c []Category
	return c, r.db.Find(&c).Error
}

func (r *repository) CreateCategory(c *Category) error {
	return r.db.Create(c).Error
}

func (r *repository) ListTags() ([]Tag, error) {
	var t []Tag
	return t, r.db.Find(&t).Error
}

func (r *repository) CreateTag(t *Tag) error {
	return r.db.Create(t).Error
}
