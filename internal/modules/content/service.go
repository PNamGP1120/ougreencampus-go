package content

import (
	"encoding/json"
	"time"
)

type Service interface {
	Create(authorID string, req CreateContentRequest) error
	GetAll() ([]ContentResponse, error)
	GetByID(id string) (ContentResponse, error)
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
	imgs, _ := json.Marshal(req.Images)

	tags, err := s.repo.FindTagsByIDs(req.TagIDs)
	if err != nil {
		return err
	}

	c := &Content{
		Title:      req.Title,
		Body:       req.Body,
		Type:       req.Type,
		CoverImage: req.CoverImage,
		Images:     imgs,
		CategoryID: req.CategoryID,
		Tags:       tags,
		AuthorID:   authorID,
	}

	return s.repo.Create(c)
}

func (s *service) GetAll() ([]ContentResponse, error) {
	list, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	res := make([]ContentResponse, 0)
	for _, c := range list {
		var imgs []string
		_ = json.Unmarshal(c.Images, &imgs)

		res = append(res, ContentResponse{
			ID:         c.ID,
			Title:      c.Title,
			Body:       c.Body,
			Type:       c.Type,
			CoverImage: c.CoverImage,
			Images:     imgs,
			IsFeatured: c.IsFeatured,
			Category:   c.Category,
			Tags:       c.Tags,
			CreatedAt:  c.CreatedAt.Format(time.RFC3339),
		})
	}

	return res, nil
}

func (s *service) GetByID(id string) (ContentResponse, error) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return ContentResponse{}, err
	}

	var imgs []string
	_ = json.Unmarshal(c.Images, &imgs)

	return ContentResponse{
		ID:         c.ID,
		Title:      c.Title,
		Body:       c.Body,
		Type:       c.Type,
		CoverImage: c.CoverImage,
		Images:     imgs,
		IsFeatured: c.IsFeatured,
		Category:   c.Category,
		Tags:       c.Tags,
		CreatedAt:  c.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *service) Update(id string, req UpdateContentRequest) error {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if req.Title != nil {
		c.Title = *req.Title
	}
	if req.Body != nil {
		c.Body = *req.Body
	}
	if req.Type != nil {
		c.Type = *req.Type
	}
	if req.CoverImage != nil {
		c.CoverImage = *req.CoverImage
	}
	if req.Images != nil {
		imgs, _ := json.Marshal(*req.Images)
		c.Images = imgs
	}
	if req.IsFeatured != nil {
		c.IsFeatured = *req.IsFeatured
	}
	if req.CategoryID != nil {
		c.CategoryID = req.CategoryID
	}
	if req.TagIDs != nil {
		tags, err := s.repo.FindTagsByIDs(*req.TagIDs)
		if err != nil {
			return err
		}
		c.Tags = tags
	}

	return s.repo.Update(c)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
