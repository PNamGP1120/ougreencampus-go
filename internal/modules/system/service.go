package system

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

// ================= SYSTEM CONFIG =================

func (s *Service) GetConfig() ([]SystemConfig, error) {
	var items []SystemConfig
	err := s.repo.db.Order("key asc").Find(&items).Error
	return items, err
}

func (s *Service) UpdateConfig(key, value string) error {
	var cfg SystemConfig
	err := s.repo.db.Where("key=?", key).First(&cfg).Error
	if err != nil {
		return s.repo.db.Create(&SystemConfig{
			Key:   key,
			Value: value,
		}).Error
	}
	return s.repo.db.Model(&cfg).Update("value", value).Error
}

// ================= REPORT =================

func (s *Service) Overview() (map[string]int64, error) {
	var users, events, activities int64

	s.repo.db.Table("users").Count(&users)
	s.repo.db.Table("events").Count(&events)
	s.repo.db.Table("activities").Count(&activities)

	return map[string]int64{
		"users":      users,
		"events":     events,
		"activities": activities,
	}, nil
}

func (s *Service) EventReport() ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	var rows []struct {
		ID    uint
		Title string
	}

	s.repo.db.Table("events").Select("id, title").Scan(&rows)
	for _, r := range rows {
		var count int64
		s.repo.db.Table("event_registrations").
			Where("event_id=?", r.ID).
			Count(&count)

		result = append(result, map[string]interface{}{
			"id":            r.ID,
			"title":         r.Title,
			"registrations": count,
		})
	}

	return result, nil
}

func (s *Service) ActivityReport() ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	var rows []struct {
		ID    uint
		Title string
	}

	s.repo.db.Table("activities").Select("id, title").Scan(&rows)
	for _, r := range rows {
		var count int64
		s.repo.db.Table("activity_participants").
			Where("activity_id=?", r.ID).
			Count(&count)

		result = append(result, map[string]interface{}{
			"id":           r.ID,
			"title":        r.Title,
			"participants": count,
		})
	}

	return result, nil
}

// ================= AUDIT =================

func (s *Service) AuditLogs(userID *uint, action string, page, limit int) ([]AuditLog, int64, error) {
	var items []AuditLog
	var total int64

	q := s.repo.db.Model(&AuditLog{})
	if userID != nil {
		q = q.Where("user_id=?", *userID)
	}
	if action != "" {
		q = q.Where("action=?", action)
	}

	q.Count(&total)
	err := q.Order("id desc").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&items).Error

	return items, total, err
}

// ================= NOTIFICATION =================

func (s *Service) Notifications(uid uint, page, limit int) ([]Notification, int64, error) {
	var items []Notification
	var total int64

	s.repo.db.Model(&Notification{}).
		Where("user_id=?", uid).
		Count(&total)

	err := s.repo.db.Where("user_id=?", uid).
		Order("id desc").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&items).Error

	return items, total, err
}

func (s *Service) ReadNotification(id, uid uint) error {
	return s.repo.db.Model(&Notification{}).
		Where("id=? AND user_id=?", id, uid).
		Update("is_read", true).Error
}
