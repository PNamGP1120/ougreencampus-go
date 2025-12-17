package activity

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

/* ========= ACTIVITY ========= */

func (r *Repository) Create(a *Activity) error {
	return r.db.Create(a).Error
}

func (r *Repository) FindByID(id string) (*Activity, error) {
	var a Activity
	err := r.db.First(&a, "id = ?", id).Error
	return &a, err
}

func (r *Repository) List(filter ListActivityFilter) ([]Activity, int64) {
	var items []Activity
	var total int64

	q := r.db.Model(&Activity{})

	if filter.Search != "" {
		q = q.Where("title ILIKE ?", "%"+filter.Search+"%")
	}
	if filter.Type != "" {
		q = q.Where("type = ?", filter.Type)
	}
	if filter.Status != "" {
		q = q.Where("status = ?", filter.Status)
	}

	q.Count(&total)

	if filter.Page > 0 && filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		q = q.Offset(offset).Limit(filter.Limit)
	}

	q.Order("created_at desc").Find(&items)

	return items, total
}

func (r *Repository) Update(id string, data map[string]any) error {
	return r.db.Model(&Activity{}).Where("id = ?", id).Updates(data).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&Activity{}, "id = ?", id).Error
}

/* ========= PARTICIPANT ========= */

func (r *Repository) Join(activityID, userID string) error {
	return r.db.Create(&Participant{
		ActivityID: activityID,
		UserID:     userID,
	}).Error
}

func (r *Repository) Leave(activityID, userID string) error {
	return r.db.Where("activity_id=? AND user_id=?", activityID, userID).
		Delete(&Participant{}).Error
}

func (r *Repository) Participants(activityID string) ([]Participant, error) {
	var items []Participant
	return items, r.db.Where("activity_id=?", activityID).Find(&items).Error
}

/* ========= SUBMISSION ========= */

func (r *Repository) CreateSubmission(s *Submission) error {
	return r.db.Create(s).Error
}

func (r *Repository) ListSubmissions(activityID, status string, page, limit int) ([]Submission, int64) {
	var items []Submission
	var total int64

	q := r.db.Model(&Submission{}).Where("activity_id=?", activityID)
	if status != "" {
		q = q.Where("status=?", status)
	}

	q.Count(&total)

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		q = q.Offset(offset).Limit(limit)
	}

	q.Find(&items)
	return items, total
}

func (r *Repository) ReviewSubmission(id string, data map[string]any) error {
	return r.db.Model(&Submission{}).Where("id=?", id).Updates(data).Error
}

/* ========= CAMPAIGN ========= */

func (r *Repository) CreateTask(t *Task) error {
	return r.db.Create(t).Error
}

func (r *Repository) ListTasks(activityID string) ([]Task, error) {
	var items []Task
	return items, r.db.Where("activity_id=?", activityID).Find(&items).Error
}

func (r *Repository) SaveProgress(p *TaskProgress) error {
	return r.db.Create(p).Error
}

/* ========= PROGRAM ========= */

func (r *Repository) AddChild(parentID, childID string) error {
	return r.db.Model(&Activity{}).
		Where("id=?", childID).
		Update("parent_id", parentID).Error
}

func (r *Repository) Children(parentID string) ([]Activity, error) {
	var items []Activity
	return items, r.db.Where("parent_id=?", parentID).Find(&items).Error
}
