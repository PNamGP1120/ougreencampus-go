package activity

import "time"

type ActivityType string
type ActivityStatus string

const (
	TypeProgram  ActivityType = "program"
	TypeContest  ActivityType = "contest"
	TypeCampaign ActivityType = "campaign"

	StatusDraft     ActivityStatus = "draft"
	StatusPublished ActivityStatus = "published"
	StatusClosed    ActivityStatus = "closed"
)

type Activity struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string
	Type        ActivityType `gorm:"index"`
	Description string       `gorm:"type:text"`
	Image       *string
	Status      ActivityStatus `gorm:"index"`
	OwnerID     string         `gorm:"index"`
	ParentID    *string        `gorm:"index"` // program → children
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Participant struct {
	ID         string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ActivityID string `gorm:"index"`
	UserID     string `gorm:"index"`
	CreatedAt  time.Time
}

type Submission struct {
	ID         string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ActivityID string `gorm:"index"`
	UserID     string `gorm:"index"`
	Content    string
	FileURL    *string
	Score      *int
	Comment    *string
	Status     string `gorm:"index"` // submitted / reviewed
	CreatedAt  time.Time
}

type Task struct {
	ID         string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ActivityID string `gorm:"index"`
	Title      string
	Target     int
}

type TaskProgress struct {
	ID     string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TaskID string `gorm:"index"`
	UserID string `gorm:"index"`
	Value  int
}
