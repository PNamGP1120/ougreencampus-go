package system

import (
	"encoding/json"
	"time"

	"gorm.io/datatypes"
)

type Service interface {
	// Config
	UpsertConfig(actorID string, req UpsertConfigRequest) error
	ListConfigs() ([]SystemConfig, error)
	DeleteConfig(key string) error

	// Audit
	Audit(actorID *string, role, action, entity, entityID string, meta map[string]interface{}, ip, ua string) error
	ListAudit(limit int) ([]AuditLog, error)

	// Reports
	Overview() (*ReportOverviewResponse, error)
}

type service struct{ repo Repository }

func NewService(repo Repository) Service { return &service{repo: repo} }

func (s *service) UpsertConfig(actorID string, req UpsertConfigRequest) error {
	now := time.Now()
	cfg := &SystemConfig{
		Key:       req.Key,
		Value:     req.Value,
		UpdatedBy: &actorID,
		UpdatedAt: now,
	}
	_ = s.Audit(&actorID, "admin", "UPSERT_CONFIG", "system_config", req.Key, map[string]interface{}{"key": req.Key}, "", "")
	return s.repo.UpsertConfig(cfg)
}

func (s *service) ListConfigs() ([]SystemConfig, error) {
	return s.repo.ListConfigs()
}

func (s *service) DeleteConfig(key string) error {
	return s.repo.DeleteConfig(key)
}

func (s *service) Audit(actorID *string, role, action, entity, entityID string, meta map[string]interface{}, ip, ua string) error {
	if meta == nil {
		meta = map[string]interface{}{}
	}
	raw, _ := json.Marshal(meta)
	a := &AuditLog{
		ActorID:   actorID,
		Role:      role,
		Action:    action,
		Entity:    entity,
		EntityID:  entityID,
		Metadata:  datatypes.JSON(raw),
		IP:        ip,
		UserAgent: ua,
		CreatedAt: time.Now(),
	}
	return s.repo.CreateAudit(a)
}

func (s *service) ListAudit(limit int) ([]AuditLog, error) {
	return s.repo.ListAudit(limit)
}

func (s *service) Overview() (*ReportOverviewResponse, error) {
	users, err := s.repo.CountTable("users")
	if err != nil {
		return nil, err
	}
	contents, err := s.repo.CountTable("contents")
	if err != nil {
		return nil, err
	}
	events, err := s.repo.CountTable("events")
	if err != nil {
		return nil, err
	}
	activities, err := s.repo.CountTable("activities")
	if err != nil {
		return nil, err
	}
	submissions, _ := s.repo.CountTable("submissions")
	registrations, _ := s.repo.CountTable("event_registrations")

	return &ReportOverviewResponse{
		Users: users, Contents: contents, Events: events, Activities: activities,
		Submissions: submissions, Registrations: registrations,
	}, nil
}
