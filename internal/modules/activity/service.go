package activity

import "errors"

type Service interface {
	Create(creatorID string, req CreateActivityRequest) error
	GetAll() ([]Activity, error)
	Submit(activityID, userID string, req SubmitRequest) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(creatorID string, req CreateActivityRequest) error {
	if req.Type != "program" && req.Type != "contest" && req.Type != "campaign" {
		return errors.New("invalid activity type")
	}

	act := &Activity{
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		CreatedBy:   creatorID,
	}
	return s.repo.Create(act)
}

func (s *service) GetAll() ([]Activity, error) {
	return s.repo.FindAll()
}

func (s *service) Submit(activityID, userID string, req SubmitRequest) error {
	sub := &Submission{
		ActivityID: activityID,
		UserID:     userID,
		Content:    req.Content,
		Status:     "submitted",
	}
	return s.repo.CreateSubmission(sub)
}
