package event

import "time"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

// ===== EVENT =====

func (s *Service) List(search string, from, to *time.Time, page, limit int) ([]Event, int64, error) {
	var items []Event
	var total int64

	q := s.repo.db.Model(&Event{})
	if search != "" {
		q = q.Where("title ILIKE ?", "%"+search+"%")
	}
	if from != nil {
		q = q.Where("start_time >= ?", *from)
	}
	if to != nil {
		q = q.Where("end_time <= ?", *to)
	}

	q.Count(&total)
	err := q.Order("id desc").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&items).Error

	return items, total, err
}

func (s *Service) Create(uid uint, req CreateEventRequest) (*Event, error) {
	e := &Event{
		Title:     req.Title,
		Image:     req.Image,
		Location:  req.Location,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		CreatedBy: uid,
	}
	return e, s.repo.db.Create(e).Error
}

func (s *Service) Get(id uint) (*Event, error) {
	var e Event
	return &e, s.repo.db.First(&e, id).Error
}

func (s *Service) Update(id uint, req UpdateEventRequest) error {
	return s.repo.db.Model(&Event{}).
		Where("id=?", id).
		Updates(map[string]interface{}{
			"title": req.Title,
			"image": req.Image,
		}).Error
}

func (s *Service) Delete(id uint) error {
	return s.repo.db.Delete(&Event{}, id).Error
}

// ===== REGISTRATION =====

func (s *Service) Register(eventID, userID uint) error {
	return s.repo.db.Create(&EventRegistration{
		EventID: eventID,
		UserID:  userID,
	}).Error
}

func (s *Service) Registrations(eventID uint, search string, page, limit int) ([]EventRegistration, int64, error) {
	var items []EventRegistration
	var total int64

	q := s.repo.db.Model(&EventRegistration{}).Where("event_id=?", eventID)
	q.Count(&total)

	err := q.Offset((page - 1) * limit).
		Limit(limit).
		Find(&items).Error

	return items, total, err
}

// ===== CHECK-IN =====

func (s *Service) Checkin(eventID, userID uint) error {
	return s.repo.db.Model(&EventRegistration{}).
		Where("event_id=? AND user_id=?", eventID, userID).
		Update("checked_in", true).Error
}

// ===== STATS =====

func (s *Service) Stats(eventID uint) (map[string]int64, error) {
	var registered, checked int64

	s.repo.db.Model(&EventRegistration{}).
		Where("event_id=?", eventID).
		Count(&registered)

	s.repo.db.Model(&EventRegistration{}).
		Where("event_id=? AND checked_in=true", eventID).
		Count(&checked)

	return map[string]int64{
		"registered": registered,
		"checked_in": checked,
	}, nil
}
