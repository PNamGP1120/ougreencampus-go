package activity

import "gorm.io/gorm"

type Repository interface {
	CreateActivity(a *Activity) error
	ListActivities() ([]Activity, error)

	CreateSubmission(s *Submission) error
	ListSubmissions(activityID string) ([]Submission, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateActivity(a *Activity) error {
	return r.db.Create(a).Error
}

func (r *repository) ListActivities() ([]Activity, error) {
	var list []Activity
	err := r.db.Find(&list).Error
	return list, err
}

func (r *repository) CreateSubmission(s *Submission) error {
	return r.db.Create(s).Error
}

func (r *repository) ListSubmissions(activityID string) ([]Submission, error) {
	var list []Submission
	err := r.db.Where("activity_id = ?", activityID).Find(&list).Error
	return list, err
}
