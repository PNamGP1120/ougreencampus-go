package event

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

/* ---------- EVENT ---------- */

func (r *Repository) CreateEvent(e *Event) error {
	return r.db.Create(e).Error
}

func (r *Repository) GetEvent(id uint) (*Event, error) {
	var e Event
	return &e, r.db.First(&e, id).Error
}

func (r *Repository) UpdateEvent(e *Event) error {
	return r.db.Save(e).Error
}

func (r *Repository) DeleteEvent(id uint) error {
	return r.db.Delete(&Event{}, id).Error
}

func (r *Repository) ListEvents() ([]Event, error) {
	var items []Event
	return items, r.db.Find(&items).Error
}

/* ---------- REGISTRATION ---------- */

func (r *Repository) Register(eventID, userID uint) error {
	return r.db.Create(&Registration{
		EventID: eventID,
		UserID:  userID,
	}).Error
}

func (r *Repository) ListRegistrations(eventID uint) ([]Registration, error) {
	var items []Registration
	return items, r.db.Where("event_id = ?", eventID).Find(&items).Error
}

func (r *Repository) Checkin(eventID, userID uint) error {
	return r.db.Model(&Registration{}).
		Where("event_id = ? AND user_id = ?", eventID, userID).
		Update("checked", true).Error
}

func (r *Repository) Stats(eventID uint) (map[string]int64, error) {
	var total, checked int64
	r.db.Model(&Registration{}).Where("event_id = ?", eventID).Count(&total)
	r.db.Model(&Registration{}).Where("event_id = ? AND checked = true", eventID).Count(&checked)
	return map[string]int64{
		"registered": total,
		"checked_in": checked,
	}, nil
}
