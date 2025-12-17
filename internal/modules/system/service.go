package system

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

/* ========= CONFIG ========= */

func (s *Service) GetConfigs() ([]SystemConfig, error) {
	return s.repo.GetConfigs()
}

func (s *Service) UpdateConfig(key, value string) error {
	return s.repo.SaveConfig(key, value)
}

/* ========= REPORT ========= */

func (s *Service) OverviewReport() map[string]any {
	return map[string]any{
		"users":      0,
		"events":     0,
		"activities": 0,
	}
}

func (s *Service) EventReport() map[string]any {
	return map[string]any{
		"total_events": 0,
		"attendees":    0,
	}
}

func (s *Service) ActivityReport() map[string]any {
	return map[string]any{
		"campaigns": 0,
		"contests":  0,
	}
}

/* ========= AUDIT ========= */

func (s *Service) AuditLogs(f AuditFilter) ([]AuditLog, int64) {
	return s.repo.ListAudit(f)
}

/* ========= NOTIFICATION ========= */

func (s *Service) Notifications(userID string, f NotificationFilter) ([]Notification, int64) {
	return s.repo.ListNotifications(userID, f)
}

func (s *Service) MarkRead(id string) error {
	return s.repo.MarkRead(id)
}
