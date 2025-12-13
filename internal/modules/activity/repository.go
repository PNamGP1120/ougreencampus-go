package activity

import "gorm.io/gorm"

type Repository interface {
	Create(activity *Activity) error
	FindAll() ([]Activity, error)
	FindByID(id string) (*Activity, error)
	CreateSubmission(sub *Submission) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(activity *Activity) error {
	return r.db.Create(activity).Error
}

func (r *repository) FindAll() ([]Activity, error) {
	var acts []Activity
	err := r.db.Order("created_at desc").Find(&acts).Error
	return acts, err
}

func (r *repository) FindByID(id string) (*Activity, error) {
	var act Activity
	err := r.db.First(&act, "id = ?", id).Error
	return &act, err
}

func (r *repository) CreateSubmission(sub *Submission) error {
	return r.db.Create(sub).Error
}
