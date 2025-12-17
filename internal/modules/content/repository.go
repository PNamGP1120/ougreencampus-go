package content

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

/* ========= CONTENT ========= */

func (r *Repository) CreateContent(c *Content) error {
	return r.db.Create(c).Error
}

func (r *Repository) GetContent(id string) (*Content, error) {
	var c Content
	err := r.db.First(&c, "id = ?", id).Error
	return &c, err
}

func (r *Repository) ListContent(filter ListContentFilter) ([]Content, int64) {
	var items []Content
	var total int64

	q := r.db.Model(&Content{})

	if filter.Search != "" {
		q = q.Where("title ILIKE ?", "%"+filter.Search+"%")
	}
	if filter.CategoryID != "" {
		q = q.Where("category_id = ?", filter.CategoryID)
	}
	if filter.TagID != "" {
		q = q.Joins(
			"JOIN content_tags ON content_tags.content_id = contents.id",
		).Where("content_tags.tag_id = ?", filter.TagID)
	}

	q.Count(&total)

	if filter.Page > 0 && filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		q = q.Offset(offset).Limit(filter.Limit)
	}

	q.Order("created_at desc").Find(&items)

	return items, total
}

func (r *Repository) UpdateContent(id string, data map[string]any) error {
	return r.db.Model(&Content{}).Where("id = ?", id).Updates(data).Error
}

func (r *Repository) DeleteContent(id string) error {
	return r.db.Delete(&Content{}, "id = ?", id).Error
}

/* ========= CATEGORY ========= */

func (r *Repository) ListCategories() ([]Category, error) {
	var items []Category
	return items, r.db.Find(&items).Error
}

func (r *Repository) CreateCategory(c *Category) error {
	return r.db.Create(c).Error
}

/* ========= TAG ========= */

func (r *Repository) ListTags() ([]Tag, error) {
	var items []Tag
	return items, r.db.Find(&items).Error
}

func (r *Repository) CreateTag(t *Tag) error {
	return r.db.Create(t).Error
}

/* ========= MEDIA ========= */

func (r *Repository) CreateMedia(m *Media) error {
	return r.db.Create(m).Error
}

func (r *Repository) ListMedia(page, limit int) ([]Media, int64) {
	var items []Media
	var total int64

	q := r.db.Model(&Media{})
	q.Count(&total)

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		q = q.Offset(offset).Limit(limit)
	}

	q.Order("created_at desc").Find(&items)
	return items, total
}

func (r *Repository) DeleteMedia(id string) error {
	return r.db.Delete(&Media{}, "id = ?", id).Error
}
