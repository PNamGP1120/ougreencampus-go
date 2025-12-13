package event

import "errors"

type Service interface {
	Create(creatorID string, req CreateEventRequest) error
	GetAll() ([]Event, error)
	Register(eventID, userID string) error
	CheckIn(eventID, userID string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(creatorID string, req CreateEventRequest) error {
	e := &Event{
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Location:    req.Location,
		Capacity:    req.Capacity,
		CreatedBy:   creatorID,
	}
	return s.repo.Create(e)
}

func (s *service) GetAll() ([]Event, error) {
	return s.repo.FindAll()
}

func (s *service) Register(eventID, userID string) error {
	count, _ := s.repo.CountRegistrations(eventID)
	event, _ := s.repo.FindByID(eventID)

	if event.Capacity > 0 && int(count) >= event.Capacity {
		return errors.New("event is full")
	}

	reg := &EventRegistration{
		EventID: eventID,
		UserID:  userID,
	}
	return s.repo.Register(reg)
}

func (s *service) CheckIn(eventID, userID string) error {
	return s.repo.CheckIn(eventID, userID)
}
