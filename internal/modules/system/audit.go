package system

import (
	"time"

	"gorm.io/gorm"
)

type AuditLog struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ActorID   string
	Action    string
	Target    string
	CreatedAt time.Time
}

type AuditRepository interface {
	Log(actorID, action, target string) error
}

type auditRepo struct {
	db *gorm.DB
}

func NewAuditRepository(db *gorm.DB) AuditRepository {
	return &auditRepo{db: db}
}

func (r *auditRepo) Log(actorID, action, target string) error {
	log := &AuditLog{
		ActorID: actorID,
		Action:  action,
		Target:  target,
	}
	return r.db.Create(log).Error
}
