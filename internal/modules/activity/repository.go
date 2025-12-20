package activity

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

/* ---------- ACTIVITY ---------- */

func (r *Repository) CreateActivity(a *Activity) error {
	return r.db.Create(a).Error
}

func (r *Repository) GetActivity(id uint) (*Activity, error) {
	var a Activity
	return &a, r.db.First(&a, id).Error
}

func (r *Repository) UpdateActivity(a *Activity) error {
	return r.db.Save(a).Error
}

func (r *Repository) DeleteActivity(id uint) error {
	return r.db.Delete(&Activity{}, id).Error
}

func (r *Repository) ListActivities() ([]Activity, error) {
	var items []Activity
	return items, r.db.Find(&items).Error
}

/* ---------- PARTICIPANT ---------- */

func (r *Repository) JoinActivity(userID, activityID uint) error {
	return r.db.Create(&Participant{UserID: userID, ActivityID: activityID}).Error
}

func (r *Repository) LeaveActivity(userID, activityID uint) error {
	return r.db.Where("user_id = ? AND activity_id = ?", userID, activityID).
		Delete(&Participant{}).Error
}

func (r *Repository) ListParticipants(activityID uint) ([]Participant, error) {
	var items []Participant
	return items, r.db.Where("activity_id = ?", activityID).Find(&items).Error
}

/* ---------- CONTEST ---------- */

func (r *Repository) Submit(s *Submission) error {
	return r.db.Create(s).Error
}

func (r *Repository) ListSubmissions(activityID uint) ([]Submission, error) {
	var items []Submission
	return items, r.db.Where("activity_id = ?", activityID).Find(&items).Error
}

func (r *Repository) ReviewSubmission(id uint, score int, comment string) error {
	return r.db.Model(&Submission{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"score":   score,
			"comment": comment,
			"status":  "reviewed",
		}).Error
}

/* ---------- CAMPAIGN ---------- */

func (r *Repository) CreateTask(t *Task) error {
	return r.db.Create(t).Error
}

func (r *Repository) ListTasks(activityID uint) ([]Task, error) {
	var items []Task
	return items, r.db.Where("activity_id = ?", activityID).Find(&items).Error
}

func (r *Repository) AddProgress(p *TaskProgress) error {
	return r.db.Create(p).Error
}

/* ---------- PROGRAM ---------- */

func (r *Repository) AddChild(parent, child uint) error {
	return r.db.Create(&ProgramChild{ParentID: parent, ChildID: child}).Error
}

func (r *Repository) ListChildren(parent uint) ([]Activity, error) {
	var items []Activity
	return items, r.db.
		Joins("JOIN program_children pc ON pc.child_id = activities.id").
		Where("pc.parent_id = ?", parent).
		Find(&items).Error
}
