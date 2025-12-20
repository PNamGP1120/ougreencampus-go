package event

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title       string
	Time        string
	Location    string
	Image       string
	OrganizerID uint
}

type Registration struct {
	gorm.Model
	EventID uint
	UserID  uint
	Checked bool
}
