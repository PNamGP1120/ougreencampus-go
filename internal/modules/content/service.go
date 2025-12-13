package content

type Service interface {
	Create(authorID string, req CreateContentRequest) error
	GetAll() ([]ContentResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(authorID string, req CreateContentRequest) error {
	content := &Content{
		Title:    req.Title,
		Body:     req.Body,
		Type:     req.Type,
		Status:   "published",
		AuthorID: authorID,
	}
	return s.repo.Create(content)
}

func (s *service) GetAll() ([]ContentResponse, error) {
	contents, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []ContentResponse
	for _, c := range contents {
		res = append(res, ContentResponse{
			ID:     c.ID,
			Title:  c.Title,
			Body:   c.Body,
			Type:   c.Type,
			Status: c.Status,
		})
	}
	return res, nil
}
