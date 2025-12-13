package event

import "gorm.io/gorm"

type Repository interface {
	Create(e *Event) error
	FindAll() ([]Event, error)
	FindByID(id string) (*Event, error)
	Update(e *Event) error

	Register(r *EventRegistration) error
	CountRegistrations(eventID string) (int64, error)
	CheckIn(eventID, userID string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(e *Event) error {
	return r.db.Create(e).Error
}

func (r *repository) FindAll() ([]Event, error) {
	var list []Event
	err := r.db.Order("start_time desc").Find(&list).Error
	return list, err
}

func (r *repository) FindByID(id string) (*Event, error) {
	var e Event
	err := r.db.First(&e, "id = ?", id).Error
	return &e, err
}

func (r *repository) Update(e *Event) error {
	return r.db.Save(e).Error
}

func (r *repository) Register(reg *EventRegistration) error {
	return r.db.Create(reg).Error
}

func (r *repository) CountRegistrations(eventID string) (int64, error) {
	var count int64
	err := r.db.Model(&EventRegistration{}).
		Where("event_id = ?", eventID).
		Count(&count).Error
	return count, err
}

func (r *repository) CheckIn(eventID, userID string) error {
	return r.db.Model(&EventRegistration{}).
		Where("event_id = ? AND user_id = ?", eventID, userID).
		Update("checked_in", true).Error
}
