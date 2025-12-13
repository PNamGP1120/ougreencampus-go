package content

import "errors"

type Service interface {
	Create(authorID string, req CreateContentRequest) error
	GetAll() ([]Content, error)
	Update(id string, req UpdateContentRequest) error
	Delete(id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(authorID string, req CreateContentRequest) error {
	c := &Content{
		Title:    req.Title,
		Body:     req.Body,
		Type:     req.Type,
		AuthorID: authorID,
	}
	return s.repo.Create(c)
}

func (s *service) GetAll() ([]Content, error) {
	return s.repo.FindAll()
}

func (s *service) Update(id string, req UpdateContentRequest) error {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("content not found")
	}

	if req.Title != "" {
		c.Title = req.Title
	}
	if req.Body != "" {
		c.Body = req.Body
	}
	if req.Type != "" {
		c.Type = req.Type
	}
	if req.IsFeatured != nil {
		c.IsFeatured = *req.IsFeatured
	}

	return s.repo.Update(c)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
