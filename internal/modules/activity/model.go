package activity

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	TypeProgram  = "program"
	TypeContest  = "contest"
	TypeCampaign = "campaign"
)

type Activity struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string
	Description string
	Type        string
	Status      string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Submission struct {
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ActivityID string
	UserID     string
	Content    string
	Status     string
	Score      *int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type CampaignTask struct {
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ActivityID string
	Title      string
	Points     int
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type CampaignProgress struct {
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ActivityID string
	TaskID     string
	UserID     string
	Status     string
	Evidence   datatypes.JSON
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
