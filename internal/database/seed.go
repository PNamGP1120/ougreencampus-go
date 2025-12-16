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
	// Guard: ensure tables exist
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

		seedUsers := []struct {
			Email  string
			Role   string
			Avatar string
		}{
			{"user1@ou.edu.vn", "admin", "https://i.pravatar.cc/150?img=1"},
			{"user2@ou.edu.vn", "organizer", "https://i.pravatar.cc/150?img=2"},
			{"user3@ou.edu.vn", "student", "https://i.pravatar.cc/150?img=3"},
			{"user4@ou.edu.vn", "student", "https://i.pravatar.cc/150?img=4"},
			{"user5@ou.edu.vn", "student", "https://i.pravatar.cc/150?img=5"},
			{"user6@ou.edu.vn", "admin", "https://i.pravatar.cc/150?img=6"},
			{"user7@ou.edu.vn", "organizer", "https://i.pravatar.cc/150?img=7"},
			{"user8@ou.edu.vn", "student", "https://i.pravatar.cc/150?img=8"},
			{"user9@ou.edu.vn", "admin", "https://i.pravatar.cc/150?img=9"},
			{"user10@ou.edu.vn", "organizer", "https://i.pravatar.cc/150?img=10"},
		}

		for _, u := range seedUsers {
			if err := db.Create(&user.User{
				Email:    u.Email,
				Password: hash,
				Role:     u.Role,
				Avatar:   u.Avatar,
				IsActive: true,
			}).Error; err != nil {
				return fmt.Errorf("seed users failed for %s: %w", u.Email, err)
			}
		}
	}

	var admin user.User
	var organizer user.User
	var student user.User

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

	// ===== CONTENT: Categories + Tags + Contents =====
	if !db.Migrator().HasTable(&content.Content{}) {
		return errors.New("seed aborted: contents table does not exist")
	}
	if !db.Migrator().HasTable(&content.Category{}) {
		return errors.New("seed aborted: categories table does not exist")
	}
	if !db.Migrator().HasTable(&content.Tag{}) {
		return errors.New("seed aborted: tags table does not exist")
	}

	// Seed categories if empty
	if err := db.Model(&content.Category{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		cats := []content.Category{
			{Name: "Tin tức", Slug: "news"},
			{Name: "Sống xanh", Slug: "green-living"},
			{Name: "Sự kiện", Slug: "events"},
			{Name: "Chia sẻ", Slug: "stories"},
		}
		for _, c := range cats {
			if err := db.Create(&c).Error; err != nil {
				return fmt.Errorf("seed categories failed: %w", err)
			}
		}
	}

	// Seed tags if empty
	if err := db.Model(&content.Tag{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		tags := []content.Tag{
			{Name: "OU", Slug: "ou"},
			{Name: "GreenCampus", Slug: "greencampus"},
			{Name: "Recycling", Slug: "recycling"},
			{Name: "Sustainability", Slug: "sustainability"},
			{Name: "Workshop", Slug: "workshop"},
		}
		for _, t := range tags {
			if err := db.Create(&t).Error; err != nil {
				return fmt.Errorf("seed tags failed: %w", err)
			}
		}
	}

	// Fetch some categories/tags for association
	var catNews content.Category
	_ = db.First(&catNews, "slug = ?", "news").Error

	var tagOU, tagGC, tagSus content.Tag
	_ = db.First(&tagOU, "slug = ?", "ou").Error
	_ = db.First(&tagGC, "slug = ?", "greencampus").Error
	_ = db.First(&tagSus, "slug = ?", "sustainability").Error

	// Seed contents if empty
	if err := db.Model(&content.Content{}).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		imgs := datatypes.JSON([]byte(`["https://picsum.photos/seed/ougc1/800/500","https://picsum.photos/seed/ougc2/800/500"]`))

		for i := 1; i <= 10; i++ {
			cid := catNews.ID
			cc := &content.Content{
				Title:      fmt.Sprintf("Content %d", i),
				Body:       "Seeded content with category + tags + images",
				Type:       "news",
				CoverImage: fmt.Sprintf("https://picsum.photos/seed/cover%d/900/600", i),
				Images:     imgs,
				IsFeatured: i%2 == 0,
				AuthorID:   adminID,
				CategoryID: &cid,
				Tags:       []content.Tag{tagOU, tagGC, tagSus},
			}
			if err := db.Create(cc).Error; err != nil {
				return fmt.Errorf("seed contents failed: %w", err)
			}
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
			if err := db.Create(&event.Event{
				Title:       fmt.Sprintf("Event %d", i),
				Description: "Seeded event",
				StartTime:   time.Now().AddDate(0, 0, i),
				EndTime:     time.Now().AddDate(0, 0, i).Add(2 * time.Hour),
				Location:    "OU Campus",
				Capacity:    50,
				CreatedBy:   organizerID,
			}).Error; err != nil {
				return fmt.Errorf("seed events failed: %w", err)
			}
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
		if err := db.Limit(10).Find(&events).Error; err != nil {
			return err
		}
		for i, e := range events {
			if err := db.Create(&event.EventRegistration{
				EventID:   e.ID,
				UserID:    studentID,
				CheckedIn: i%2 == 0,
			}).Error; err != nil {
				return fmt.Errorf("seed registrations failed: %w", err)
			}
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
			if err := db.Create(&activity.Activity{
				Name:      fmt.Sprintf("Activity %d", i),
				Type:      types[i%len(types)],
				Status:    "published",
				CreatedBy: organizerID,
			}).Error; err != nil {
				return fmt.Errorf("seed activities failed: %w", err)
			}
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
		if err := db.Where("type = ?", activity.TypeContest).Limit(10).Find(&contests).Error; err != nil {
			return err
		}
		for _, a := range contests {
			if err := db.Create(&activity.Submission{
				ActivityID: a.ID,
				UserID:     studentID,
				Content:    "Seeded submission",
				Status:     "approved",
			}).Error; err != nil {
				return fmt.Errorf("seed submissions failed: %w", err)
			}
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
		if err := db.Where("type = ?", activity.TypeCampaign).Limit(5).Find(&campaigns).Error; err != nil {
			return err
		}
		for _, a := range campaigns {
			for i := 1; i <= 3; i++ {
				if err := db.Create(&activity.CampaignTask{
					ActivityID: a.ID,
					Title:      fmt.Sprintf("Task %d", i),
					Points:     10 * i,
					IsActive:   true,
				}).Error; err != nil {
					return fmt.Errorf("seed campaign_tasks failed: %w", err)
				}
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
		if err := db.Limit(10).Find(&tasks).Error; err != nil {
			return err
		}
		raw := datatypes.JSON([]byte(`{"proof":"seed"}`))
		for _, t := range tasks {
			if err := db.Create(&activity.CampaignProgress{
				ActivityID: t.ActivityID,
				TaskID:     t.ID,
				UserID:     studentID,
				Status:     "approved",
				Evidence:   raw,
			}).Error; err != nil {
				return fmt.Errorf("seed campaign_progress failed: %w", err)
			}
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
			if err := db.Create(&system.SystemConfig{
				Key:       fmt.Sprintf("config_%d", i),
				Value:     "value",
				UpdatedBy: &adminID,
				UpdatedAt: time.Now(),
			}).Error; err != nil {
				return fmt.Errorf("seed system_configs failed: %w", err)
			}
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
			if err := db.Create(&system.AuditLog{
				ActorID:   &adminID,
				Role:      "admin",
				Action:    "SEED",
				Entity:    "system",
				EntityID:  fmt.Sprintf("seed_%d", i),
				CreatedAt: time.Now(),
			}).Error; err != nil {
				return fmt.Errorf("seed audit_logs failed: %w", err)
			}
		}
	}

	return nil
}
