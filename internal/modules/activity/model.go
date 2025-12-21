package activity

import "time"

type Activity struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"` // program | contest | campaign
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Status      string    `json:"status"` // active | closed
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ActivityParticipant struct {
	ID         uint `gorm:"primaryKey"`
	ActivityID uint
	UserID     uint
	CreatedAt  time.Time
}

type ContestSubmission struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ActivityID uint      `json:"activity_id"`
	UserID     uint      `json:"user_id"`
	Content    string    `json:"content"`
	FileURL    string    `json:"file_url"`
	Score      *int      `json:"score"`
	Comment    *string   `json:"comment"`
	Status     string    `json:"status"` // submitted | reviewed
	CreatedAt  time.Time `json:"created_at"`
}

type CampaignTask struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ActivityID uint   `json:"activity_id"`
	Title      string `json:"title"`
	Target     int    `json:"target"`
}

type CampaignProgress struct {
	ID     uint `gorm:"primaryKey"`
	TaskID uint
	UserID uint
	Value  int
}

type ProgramRelation struct {
	ID       uint `gorm:"primaryKey"`
	ParentID uint
	ChildID  uint
}
