package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/PNamGP1120/ougreencampus-go/internal/utils"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// Guard: ensure tables exist (quick check)
	if !db.Migrator().HasTable(&user.User{}) {
		return errors.New("seed aborted: users table does not exist (run migrate first)")
	}

	// ===== USERS =====
	var count int64
	if err := db.Model(&user.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		hash, err := utils.HashPassword("123456")
		if err != nil {
			return err
		}

		// deterministic roles for easier testing
		seedUsers := []struct {
			Email string
			Role  string
		}{
			{"user1@ou.edu.vn", "admin"},
			{"user2@ou.edu.vn", "organizer"},
			{"user3@ou.edu.vn", "student"},
			{"user4@ou.edu.vn", "student"},
			{"user5@ou.edu.vn", "student"},
			{"user6@ou.edu.vn", "admin"},
			{"user7@ou.edu.vn", "organizer"},
			{"user8@ou.edu.vn", "student"},
			{"user9@ou.edu.vn", "admin"},
			{"user10@ou.edu.vn", "organizer"},
		}

		for _, u := range seedUsers {
			_ = db.Create(&user.User{
				Email:    u.Email,
				Password: hash,
				Role:     u.Role,
				IsActive: true,
			}).Error
		}
	}

	var admin user.User
	var organizer user.User
	var student user.User

	// Ensure we can find at least one of each role
	if err := db.First(&admin, "role = ? AND is_active = true", "admin").Error; err != nil {
		return fmt.Errorf("seed aborted: cannot find admin user: %w", err)
	}
	if err := db.First(&organizer, "role = ? AND is_active = true", "organizer").Error; err != nil {
		return fmt.Errorf("seed aborted: cannot find organizer user: %w", err)
	}
	if err := db.First(&student, "role = ? AND is_active = true", "student").Error; err != nil {
		return fmt.Errorf("seed aborted: cannot find student user: %w", err)
	}

	adminID := admin.ID
	organizerID := organizer.ID
	studentID := student.ID

	// ===== CONTENT =====
	if !db.Migrator().HasTable(&content.Content{}) {
		return errors.New("seed aborted: contents table does not exist")
	}
	if err := db.Model(&content.Content{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		for i := 1; i <= 10; i++ {
			_ = db.Create(&content.Content{
				Title:      fmt.Sprintf("Content %d", i),
				Body:       "Seeded content",
				Type:       "news",
				AuthorID:   adminID,
				IsFeatured: i%2 == 0,
			}).Error
		}
	}

	// ===== EVENTS =====
	if !db.Migrator().HasTable(&event.Event{}) {
		return errors.New("seed aborted: events table does not exist")
	}
	if err := db.Model(&event.Event{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		for i := 1; i <= 10; i++ {
			_ = db.Create(&event.Event{
				Title:       fmt.Sprintf("Event %d", i),
				Description: "Seeded event",
				StartTime:   time.Now().AddDate(0, 0, i),
				EndTime:     time.Now().AddDate(0, 0, i).Add(2 * time.Hour),
				Location:    "OU Campus",
				Capacity:    50,
				CreatedBy:   organizerID,
			}).Error
		}
	}

	// ===== EVENT REGISTRATIONS =====
	if !db.Migrator().HasTable(&event.EventRegistration{}) {
		return errors.New("seed aborted: event_registrations table does not exist")
	}
	if err := db.Model(&event.EventRegistration{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		var events []event.Event
		_ = db.Limit(10).Find(&events).Error
		for i, e := range events {
			_ = db.Create(&event.EventRegistration{
				EventID:   e.ID,
				UserID:    studentID,
				CheckedIn: i%2 == 0,
			}).Error
		}
	}

	// ===== ACTIVITIES =====
	if !db.Migrator().HasTable(&activity.Activity{}) {
		return errors.New("seed aborted: activities table does not exist")
	}
	if err := db.Model(&activity.Activity{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		types := []string{activity.TypeProgram, activity.TypeContest, activity.TypeCampaign}
		for i := 1; i <= 10; i++ {
			_ = db.Create(&activity.Activity{
				Name:      fmt.Sprintf("Activity %d", i),
				Type:      types[i%len(types)],
				Status:    "published",
				CreatedBy: organizerID,
			}).Error
		}
	}

	// ===== SUBMISSIONS (contest only) =====
	if !db.Migrator().HasTable(&activity.Submission{}) {
		return errors.New("seed aborted: submissions table does not exist")
	}
	if err := db.Model(&activity.Submission{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		var contests []activity.Activity
		_ = db.Where("type = ?", activity.TypeContest).Limit(10).Find(&contests).Error
		for _, a := range contests {
			_ = db.Create(&activity.Submission{
				ActivityID: a.ID,
				UserID:     studentID,
				Content:    "Seeded submission",
				Status:     "approved",
			}).Error
		}
	}

	// ===== CAMPAIGN TASKS =====
	if !db.Migrator().HasTable(&activity.CampaignTask{}) {
		return errors.New("seed aborted: campaign_tasks table does not exist")
	}
	if err := db.Model(&activity.CampaignTask{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		var campaigns []activity.Activity
		_ = db.Where("type = ?", activity.TypeCampaign).Limit(5).Find(&campaigns).Error
		for _, a := range campaigns {
			for i := 1; i <= 3; i++ {
				_ = db.Create(&activity.CampaignTask{
					ActivityID: a.ID,
					Title:      fmt.Sprintf("Task %d", i),
					Points:     10 * i,
					IsActive:   true,
				}).Error
			}
		}
	}

	// ===== CAMPAIGN PROGRESS =====
	if !db.Migrator().HasTable(&activity.CampaignProgress{}) {
		return errors.New("seed aborted: campaign_progresses table does not exist")
	}
	if err := db.Model(&activity.CampaignProgress{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		var tasks []activity.CampaignTask
		_ = db.Limit(10).Find(&tasks).Error
		raw := datatypes.JSON([]byte(`{"proof":"seed"}`))
		for _, t := range tasks {
			_ = db.Create(&activity.CampaignProgress{
				ActivityID: t.ActivityID,
				TaskID:     t.ID,
				UserID:     studentID,
				Status:     "approved",
				Evidence:   raw,
			}).Error
		}
	}

	// ===== SYSTEM CONFIG =====
	if !db.Migrator().HasTable(&system.SystemConfig{}) {
		return errors.New("seed aborted: system_configs table does not exist")
	}
	if err := db.Model(&system.SystemConfig{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		for i := 1; i <= 10; i++ {
			_ = db.Create(&system.SystemConfig{
				Key:       fmt.Sprintf("config_%d", i),
				Value:     "value",
				UpdatedBy: &adminID,
				UpdatedAt: time.Now(),
			}).Error
		}
	}

	// ===== AUDIT LOG =====
	if !db.Migrator().HasTable(&system.AuditLog{}) {
		return errors.New("seed aborted: audit_logs table does not exist")
	}
	if err := db.Model(&system.AuditLog{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		for i := 1; i <= 10; i++ {
			_ = db.Create(&system.AuditLog{
				ActorID:   &adminID,
				Role:      "admin",
				Action:    "SEED",
				Entity:    "system",
				EntityID:  fmt.Sprintf("seed_%d", i),
				CreatedAt: time.Now(),
			}).Error
		}
	}

	return nil
}
