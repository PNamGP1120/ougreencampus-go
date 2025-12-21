package activity

import "time"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

// ================= ACTIVITY =================

func (s *Service) List(search, typ, status string, from, to *time.Time, page, limit int) ([]Activity, int64, error) {
	var items []Activity
	var total int64

	q := s.repo.db.Model(&Activity{})
	if search != "" {
		q = q.Where("title ILIKE ?", "%"+search+"%")
	}
	if typ != "" {
		q = q.Where("type=?", typ)
	}
	if status != "" {
		q = q.Where("status=?", status)
	}
	if from != nil {
		q = q.Where("created_at >= ?", *from)
	}
	if to != nil {
		q = q.Where("created_at <= ?", *to)
	}

	q.Count(&total)
	err := q.Order("id desc").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&items).Error

	return items, total, err
}

func (s *Service) Create(uid uint, req CreateActivityRequest) (*Activity, error) {
	a := &Activity{
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
		Image:       req.Image,
		Status:      "active",
		CreatedBy:   uid,
	}
	return a, s.repo.db.Create(a).Error
}

func (s *Service) Get(id uint) (*Activity, error) {
	var a Activity
	return &a, s.repo.db.First(&a, id).Error
}

func (s *Service) Update(id uint, req UpdateActivityRequest) error {
	return s.repo.db.Model(&Activity{}).
		Where("id=?", id).
		Updates(map[string]interface{}{
			"title": req.Title,
			"image": req.Image,
		}).Error
}

func (s *Service) Delete(id uint) error {
	return s.repo.db.Delete(&Activity{}, id).Error
}

// ================= PARTICIPANT =================

func (s *Service) Join(aid, uid uint) error {
	return s.repo.db.Create(&ActivityParticipant{
		ActivityID: aid,
		UserID:     uid,
	}).Error
}

func (s *Service) Leave(aid, uid uint) error {
	return s.repo.db.Where("activity_id=? AND user_id=?", aid, uid).
		Delete(&ActivityParticipant{}).Error
}

func (s *Service) Participants(aid uint, page, limit int) ([]ActivityParticipant, int64, error) {
	var items []ActivityParticipant
	var total int64

	s.repo.db.Model(&ActivityParticipant{}).
		Where("activity_id=?", aid).
		Count(&total)

	err := s.repo.db.Where("activity_id=?", aid).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&items).Error

	return items, total, err
}

// ================= CONTEST =================

func (s *Service) Submit(aid, uid uint, req SubmitContestRequest) (*ContestSubmission, error) {
	sub := &ContestSubmission{
		ActivityID: aid,
		UserID:     uid,
		Content:    req.Content,
		FileURL:    req.FileURL,
		Status:     "submitted",
	}
	return sub, s.repo.db.Create(sub).Error
}

func (s *Service) Submissions(aid uint, status string, page, limit int) ([]ContestSubmission, int64, error) {
	var items []ContestSubmission
	var total int64

	q := s.repo.db.Model(&ContestSubmission{}).Where("activity_id=?", aid)
	if status != "" {
		q = q.Where("status=?", status)
	}

	q.Count(&total)
	err := q.Offset((page - 1) * limit).
		Limit(limit).
		Find(&items).Error

	return items, total, err
}

func (s *Service) Review(subID uint, score int, comment string) error {
	return s.repo.db.Model(&ContestSubmission{}).
		Where("id=?", subID).
		Updates(map[string]interface{}{
			"score":   score,
			"comment": comment,
			"status":  "reviewed",
		}).Error
}

// ================= CAMPAIGN =================

func (s *Service) CreateTask(aid uint, req CreateTaskRequest) (*CampaignTask, error) {
	t := &CampaignTask{
		ActivityID: aid,
		Title:      req.Title,
		Target:     req.Target,
	}
	return t, s.repo.db.Create(t).Error
}

func (s *Service) Tasks(aid uint, page, limit int) ([]CampaignTask, int64, error) {
	var items []CampaignTask
	var total int64

	s.repo.db.Model(&CampaignTask{}).
		Where("activity_id=?", aid).
		Count(&total)

	err := s.repo.db.Where("activity_id=?", aid).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&items).Error

	return items, total, err
}

func (s *Service) Progress(tid, uid uint, value int) error {
	return s.repo.db.Create(&CampaignProgress{
		TaskID: tid,
		UserID: uid,
		Value:  value,
	}).Error
}

func (s *Service) Metrics(aid uint) (map[string]int64, error) {
	var tasks, progress int64
	s.repo.db.Model(&CampaignTask{}).Where("activity_id=?", aid).Count(&tasks)
	s.repo.db.Model(&CampaignProgress{}).
		Joins("JOIN campaign_tasks ON campaign_tasks.id = campaign_progresses.task_id").
		Where("campaign_tasks.activity_id=?", aid).
		Count(&progress)

	return map[string]int64{
		"tasks":    tasks,
		"progress": progress,
	}, nil
}

// ================= PROGRAM =================

func (s *Service) AddChild(pid, cid uint) error {
	return s.repo.db.Create(&ProgramRelation{
		ParentID: pid,
		ChildID:  cid,
	}).Error
}

func (s *Service) Children(pid uint) ([]Activity, error) {
	var items []Activity
	err := s.repo.db.
		Joins("JOIN program_relations ON program_relations.child_id = activities.id").
		Where("program_relations.parent_id=?", pid).
		Find(&items).Error
	return items, err
}
