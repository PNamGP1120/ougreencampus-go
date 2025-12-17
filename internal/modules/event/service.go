package event

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

/* ========= EVENT ========= */

func (s *Service) CreateEvent(req CreateEventRequest, creatorID string) (string, error) {
	e := &Event{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Location:    req.Location,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Capacity:    req.Capacity,
		CreatedBy:   creatorID,
	}
	return e.ID, s.repo.Create(e)
}

func (s *Service) GetEvent(id string) (*Event, error) {
	return s.repo.FindByID(id)
}

func (s *Service) ListEvents(filter ListEventFilter) ([]Event, int64) {
	return s.repo.List(filter)
}

func (s *Service) UpdateEvent(id string, req UpdateEventRequest) error {
	return s.repo.Update(id, map[string]any{
		"title":       req.Title,
		"description": req.Description,
		"image":       req.Image,
		"location":    req.Location,
		"start_time":  req.StartTime,
		"end_time":    req.EndTime,
		"capacity":    req.Capacity,
	})
}

func (s *Service) DeleteEvent(id string) error {
	return s.repo.Delete(id)
}

/* ========= REGISTRATION ========= */

func (s *Service) Register(eventID, userID string) error {
	return s.repo.Register(eventID, userID)
}

func (s *Service) Registrations(eventID string) ([]Registration, error) {
	return s.repo.Registrations(eventID)
}

func (s *Service) Checkin(eventID, userID string) error {
	return s.repo.Checkin(eventID, userID)
}
