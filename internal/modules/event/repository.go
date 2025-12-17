package event

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

/* ========= EVENT ========= */

func (r *Repository) Create(e *Event) error {
	return r.db.Create(e).Error
}

func (r *Repository) FindByID(id string) (*Event, error) {
	var e Event
	err := r.db.First(&e, "id = ?", id).Error
	return &e, err
}

func (r *Repository) List(filter ListEventFilter) ([]Event, int64) {
	var items []Event
	var total int64

	q := r.db.Model(&Event{})

	if filter.Search != "" {
		q = q.Where("title ILIKE ?", "%"+filter.Search+"%")
	}
	if filter.From != "" {
		q = q.Where("start_time >= ?", filter.From)
	}
	if filter.To != "" {
		q = q.Where("end_time <= ?", filter.To)
	}

	q.Count(&total)

	if filter.Page > 0 && filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		q = q.Offset(offset).Limit(filter.Limit)
	}

	q.Order("start_time desc").Find(&items)
	return items, total
}

func (r *Repository) Update(id string, data map[string]any) error {
	return r.db.Model(&Event{}).Where("id = ?", id).Updates(data).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Event{}, "id = ?", id).Error
}

/* ========= REGISTRATION ========= */

func (r *Repository) Register(eventID, userID string) error {
	return r.db.Create(&Registration{
		EventID: eventID,
		UserID:  userID,
	}).Error
}

func (r *Repository) Registrations(eventID string) ([]Registration, error) {
	var items []Registration
	return items, r.db.Where("event_id = ?", eventID).Find(&items).Error
}

func (r *Repository) Checkin(eventID, userID string) error {
	return r.db.Model(&Registration{}).
		Where("event_id=? AND user_id=?", eventID, userID).
		Update("checked_in", true).Error
}
