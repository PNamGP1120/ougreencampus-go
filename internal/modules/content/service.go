package content

import "errors"

type Service interface {
	List(filter map[string]interface{}, page, limit int) ([]Content, int64, error)
	Create(authorID uint, req CreateContentRequest) (*Content, error)
	Get(id uint) (*Content, error)
	Update(id uint, req UpdateContentRequest) error
	Delete(id uint) error

	Categories() ([]Category, error)
	CreateCategory(name string) (*Category, error)

	Tags() ([]Tag, error)
	CreateTag(name string) (*Tag, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) List(filter map[string]interface{}, page, limit int) ([]Content, int64, error) {
	return s.repo.List(filter, page, limit)
}

func (s *service) Create(authorID uint, req CreateContentRequest) (*Content, error) {
	c := &Content{
		Title:      req.Title,
		Body:       req.Body,
		Image:      req.Image,
		CategoryID: req.CategoryID,
		AuthorID:   authorID,
	}

	if err := s.repo.Create(c); err != nil {
		return nil, err
	}

	if len(req.Tags) > 0 {
		var tags []Tag
		for _, id := range req.Tags {
			tags = append(tags, Tag{ID: id})
		}
		c.Tags = tags
		s.repo.Update(c)
	}

	return c, nil
}

func (s *service) Get(id uint) (*Content, error) {
	return s.repo.FindByID(id)
}

func (s *service) Update(id uint, req UpdateContentRequest) error {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	c.Title = req.Title
	c.Body = req.Body
	c.Image = req.Image
	return s.repo.Update(c)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *service) Categories() ([]Category, error) {
	return s.repo.ListCategories()
}

func (s *service) CreateCategory(name string) (*Category, error) {
	if name == "" {
		return nil, errors.New("name required")
	}
	c := &Category{Name: name}
	return c, s.repo.CreateCategory(c)
}

func (s *service) Tags() ([]Tag, error) {
	return s.repo.ListTags()
}

func (s *service) CreateTag(name string) (*Tag, error) {
	if name == "" {
		return nil, errors.New("name required")
	}
	t := &Tag{Name: name}
	return t, s.repo.CreateTag(t)
}
