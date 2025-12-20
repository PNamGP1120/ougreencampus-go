package activity

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Title       string
	Type        string // program | contest | campaign
	Description string `gorm:"type:text"`
	Image       string
	OrganizerID uint
}

type Participant struct {
	gorm.Model
	UserID     uint
	ActivityID uint
}

type Submission struct {
	gorm.Model
	UserID     uint
	ActivityID uint
	Content    string
	FileURL    string
	Status     string // pending | reviewed
	Score      int
	Comment    string
}

type Task struct {
	gorm.Model
	ActivityID uint
	Title      string
	Target     int
}

type TaskProgress struct {
	gorm.Model
	TaskID uint
	UserID uint
	Value  int
}

type ProgramChild struct {
	gorm.Model
	ParentID uint
	ChildID  uint
}
