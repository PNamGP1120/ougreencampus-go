package event

type Service interface {
	Create(creatorID string, req CreateEventRequest) error
	GetAll() ([]Event, error)
	Register(eventID, userID string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(creatorID string, req CreateEventRequest) error {
	event := &Event{
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Location:    req.Location,
		Capacity:    req.Capacity,
		CreatedBy:   creatorID,
	}
	return s.repo.Create(event)
}

func (s *service) GetAll() ([]Event, error) {
	return s.repo.FindAll()
}

func (s *service) Register(eventID, userID string) error {
	return s.repo.Register(eventID, userID)
}
