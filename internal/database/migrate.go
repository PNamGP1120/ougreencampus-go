package database

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	// IMPORTANT: order does not matter for GORM basic, but keep stable
	return db.AutoMigrate(
		&user.User{},

		&content.Content{},

		&event.Event{},
		&event.EventRegistration{},

		&activity.Activity{},
		&activity.Submission{},
		&activity.CampaignTask{},
		&activity.CampaignProgress{},

		&system.SystemConfig{},
		&system.AuditLog{},
	)
}
