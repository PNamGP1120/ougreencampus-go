package activity

import "errors"

type Service interface {
	CreateActivity(name, typ, creator string) error
	ListActivities() ([]Activity, error)
	SubmitContest(activityID, userID, content string) error
	ListSubmissions(activityID string) ([]Submission, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateActivity(name, typ, creator string) error {
	if typ != TypeProgram && typ != TypeContest && typ != TypeCampaign {
		return errors.New("invalid activity type")
	}

	return s.repo.CreateActivity(&Activity{
		Name:      name,
		Type:      typ,
		Status:    "published",
		CreatedBy: creator,
	})
}

func (s *service) ListActivities() ([]Activity, error) {
	return s.repo.ListActivities()
}

func (s *service) SubmitContest(activityID, userID, content string) error {
	return s.repo.CreateSubmission(&Submission{
		ActivityID: activityID,
		UserID:     userID,
		Content:    content,
		Status:     "pending",
	})
}

func (s *service) ListSubmissions(activityID string) ([]Submission, error) {
	return s.repo.ListSubmissions(activityID)
}
