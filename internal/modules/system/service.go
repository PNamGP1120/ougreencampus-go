package system

import "time"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{r}
}

/* ---------- REPORT ---------- */

func (s *Service) OverviewReport(from, to string) map[string]interface{} {
	return map[string]interface{}{
		"from":       from,
		"to":         to,
		"users":      0,
		"events":     0,
		"activities": 0,
	}
}

func (s *Service) EventReport(from, to string) map[string]interface{} {
	return map[string]interface{}{
		"type":         "event",
		"from":         from,
		"to":           to,
		"generated_at": time.Now(),
	}
}

func (s *Service) ActivityReport(from, to string) map[string]interface{} {
	return map[string]interface{}{
		"type":         "activity",
		"from":         from,
		"to":           to,
		"generated_at": time.Now(),
	}
}
