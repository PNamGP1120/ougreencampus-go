package content

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

/* ---------- CONTENT ---------- */

func (r *Repository) CreateContent(c *Content) error {
	return r.db.Create(c).Error
}

func (r *Repository) GetContent(id uint) (*Content, error) {
	var c Content
	err := r.db.Preload("Category").Preload("Tags").First(&c, id).Error
	return &c, err
}

func (r *Repository) UpdateContent(c *Content) error {
	return r.db.Save(c).Error
}

func (r *Repository) DeleteContent(id uint) error {
	return r.db.Delete(&Content{}, id).Error
}

func (r *Repository) ListContents() ([]Content, error) {
	var items []Content
	return items, r.db.Preload("Category").Preload("Tags").Find(&items).Error
}

/* ---------- CATEGORY ---------- */

func (r *Repository) ListCategories() ([]Category, error) {
	var c []Category
	return c, r.db.Find(&c).Error
}

func (r *Repository) CreateCategory(c *Category) error {
	return r.db.Create(c).Error
}

/* ---------- TAG ---------- */

func (r *Repository) ListTags() ([]Tag, error) {
	var t []Tag
	return t, r.db.Find(&t).Error
}

func (r *Repository) CreateTag(t *Tag) error {
	return r.db.Create(t).Error
}

/* ---------- MEDIA ---------- */

func (r *Repository) CreateMedia(m *Media) error {
	return r.db.Create(m).Error
}

func (r *Repository) ListMedia() ([]Media, error) {
	var m []Media
	return m, r.db.Find(&m).Error
}

func (r *Repository) DeleteMedia(id uint) error {
	return r.db.Delete(&Media{}, id).Error
}

func (r *Repository) AttachMedia(id uint, refType string, refID uint) error {
	return r.db.Model(&Media{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"ref_type": refType,
			"ref_id":   refID,
		}).Error
}
