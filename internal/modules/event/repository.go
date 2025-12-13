package event

import "gorm.io/gorm"

type Repository interface {
	Create(event *Event) error
	FindAll() ([]Event, error)
	Register(eventID, userID string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(event *Event) error {
	return r.db.Create(event).Error
}

func (r *repository) FindAll() ([]Event, error) {
	var events []Event
	err := r.db.Order("start_time asc").Find(&events).Error
	return events, err
}

func (r *repository) Register(eventID, userID string) error {
	reg := &EventRegistration{
		EventID: eventID,
		UserID:  userID,
	}
	return r.db.Create(reg).Error
}
