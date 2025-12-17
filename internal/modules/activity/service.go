package activity

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

/* ========= ACTIVITY ========= */

func (s *Service) CreateActivity(req CreateActivityRequest, ownerID string) (string, error) {
	a := &Activity{
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
		Image:       req.Image,
		Status:      StatusDraft,
		OwnerID:     ownerID,
	}
	return a.ID, s.repo.Create(a)
}

func (s *Service) GetActivity(id string) (*Activity, error) {
	return s.repo.FindByID(id)
}

func (s *Service) ListActivities(filter ListActivityFilter) ([]Activity, int64) {
	return s.repo.List(filter)
}

func (s *Service) UpdateActivity(id string, req UpdateActivityRequest) error {
	return s.repo.Update(id, map[string]any{
		"title":       req.Title,
		"description": req.Description,
		"image":       req.Image,
		"status":      req.Status,
	})
}

func (s *Service) DeleteActivity(id string) error {
	return s.repo.Delete(id)
}

/* ========= PARTICIPANT ========= */

func (s *Service) Join(activityID, userID string) error {
	return s.repo.Join(activityID, userID)
}

func (s *Service) Leave(activityID, userID string) error {
	return s.repo.Leave(activityID, userID)
}

func (s *Service) Participants(activityID string) ([]Participant, error) {
	return s.repo.Participants(activityID)
}

/* ========= SUBMISSION ========= */

func (s *Service) Submit(activityID, userID string, req SubmitRequest) (string, error) {
	sub := &Submission{
		ActivityID: activityID,
		UserID:     userID,
		Content:    req.Content,
		FileURL:    req.FileURL,
		Status:     "submitted",
	}
	return sub.ID, s.repo.CreateSubmission(sub)
}

func (s *Service) Submissions(activityID, status string, page, limit int) ([]Submission, int64) {
	return s.repo.ListSubmissions(activityID, status, page, limit)
}

func (s *Service) ReviewSubmission(id string, req ReviewSubmissionRequest) error {
	return s.repo.ReviewSubmission(id, map[string]any{
		"score":   req.Score,
		"comment": req.Comment,
		"status":  "reviewed",
	})
}

/* ========= CAMPAIGN ========= */

func (s *Service) CreateTask(activityID string, req CreateTaskRequest) (string, error) {
	t := &Task{
		ActivityID: activityID,
		Title:      req.Title,
		Target:     req.Target,
	}
	return t.ID, s.repo.CreateTask(t)
}

func (s *Service) Tasks(activityID string) ([]Task, error) {
	return s.repo.ListTasks(activityID)
}

func (s *Service) Progress(taskID, userID string, value int) error {
	p := &TaskProgress{
		TaskID: taskID,
		UserID: userID,
		Value:  value,
	}
	return s.repo.SaveProgress(p)
}

/* ========= PROGRAM ========= */

func (s *Service) AddChild(parentID, childID string) error {
	return s.repo.AddChild(parentID, childID)
}

func (s *Service) Children(parentID string) ([]Activity, error) {
	return s.repo.Children(parentID)
}
