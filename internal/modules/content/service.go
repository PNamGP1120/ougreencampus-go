package content

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

/* ========= CONTENT ========= */

func (s *Service) CreateContent(req CreateContentRequest, authorID string) (string, error) {
	c := &Content{
		Title:      req.Title,
		Body:       req.Body,
		Image:      req.Image,
		AuthorID:   authorID,
		CategoryID: req.CategoryID,
	}
	return c.ID, s.repo.CreateContent(c)
}

func (s *Service) GetContent(id string) (*Content, error) {
	return s.repo.GetContent(id)
}

func (s *Service) ListContent(filter ListContentFilter) ([]Content, int64) {
	return s.repo.ListContent(filter)
}

func (s *Service) UpdateContent(id string, req UpdateContentRequest) error {
	return s.repo.UpdateContent(id, map[string]any{
		"title": req.Title,
		"body":  req.Body,
		"image": req.Image,
	})
}

func (s *Service) DeleteContent(id string) error {
	return s.repo.DeleteContent(id)
}

/* ========= CATEGORY ========= */

func (s *Service) ListCategories() ([]Category, error) {
	return s.repo.ListCategories()
}

func (s *Service) CreateCategory(name string) (string, error) {
	c := &Category{Name: name}
	return c.ID, s.repo.CreateCategory(c)
}

/* ========= TAG ========= */

func (s *Service) ListTags() ([]Tag, error) {
	return s.repo.ListTags()
}

func (s *Service) CreateTag(name string) (string, error) {
	t := &Tag{Name: name}
	return t.ID, s.repo.CreateTag(t)
}

/* ========= MEDIA ========= */

func (s *Service) UploadMedia() (string, string) {
	m := &Media{
		URL:  "https://cdn.ougreencampus/media/sample.jpg",
		Type: "image",
	}
	_ = s.repo.CreateMedia(m)
	return m.ID, m.URL
}

func (s *Service) ListMedia(page, limit int) ([]Media, int64) {
	return s.repo.ListMedia(page, limit)
}

func (s *Service) DeleteMedia(id string) error {
	return s.repo.DeleteMedia(id)
}

func (s *Service) AttachMedia(id, typ, refID string) error {
	// placeholder
	return nil
}
