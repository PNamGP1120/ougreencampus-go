package database

import (
	"fmt"
	"time"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/PNamGP1120/ougreencampus-go/internal/utils"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	fmt.Println("🌱 Seeding database...")

	// ================= USERS =================
	var count int64
	db.Model(&user.User{}).Count(&count)
	if count == 0 {
		roles := []string{
			config.RoleAdmin,
			config.RoleOrganizer,
			config.RoleStudent,
			config.RoleStudent,
			config.RoleStudent,
		}

		for i := 1; i <= 10; i++ {
			role := roles[i%len(roles)]
			password, _ := utils.HashPassword("123456")

			u := user.User{
				Email:    fmt.Sprintf("user%d@ou.edu.vn", i),
				Password: password,
				Role:     role,
				IsActive: true,
			}
			db.Create(&u)
		}
	}

	// Lấy 1 admin & 1 organizer để gán quan hệ
	var admin user.User
	var organizer user.User
	db.Where("role = ?", config.RoleAdmin).First(&admin)
	db.Where("role = ?", config.RoleOrganizer).First(&organizer)

	// ================= CONTENT =================
	db.Model(&content.Content{}).Count(&count)
	if count == 0 {
		types := []string{"news", "blog", "green_news"}

		for i := 1; i <= 10; i++ {
			c := content.Content{
				Title:    fmt.Sprintf("Green Content %d", i),
				Body:     "This is sample green campus content.",
				Type:     types[i%len(types)],
				Status:   "published",
				AuthorID: admin.ID,
			}
			db.Create(&c)
		}
	}

	// ================= EVENTS =================
	db.Model(&event.Event{}).Count(&count)
	if count == 0 {
		for i := 1; i <= 10; i++ {
			e := event.Event{
				Title:       fmt.Sprintf("Green Event %d", i),
				Description: "OU GreenCampus event",
				StartTime:   time.Now().AddDate(0, 0, i),
				EndTime:     time.Now().AddDate(0, 0, i+1),
				Location:    "OU Campus",
				Capacity:    100,
				CreatedBy:   organizer.ID,
			}
			db.Create(&e)
		}
	}

	// ================= EVENT REGISTRATION =================
	db.Model(&event.EventRegistration{}).Count(&count)
	if count == 0 {
		var events []event.Event
		var students []user.User

		db.Find(&events)
		db.Where("role = ?", config.RoleStudent).Find(&students)

		for i := 0; i < 10 && i < len(events) && i < len(students); i++ {
			r := event.EventRegistration{
				EventID: events[i].ID,
				UserID:  students[i].ID,
			}
			db.Create(&r)
		}
	}

	// ================= ACTIVITIES =================
	db.Model(&activity.Activity{}).Count(&count)
	if count == 0 {
		types := []string{"program", "contest", "campaign"}

		for i := 1; i <= 10; i++ {
			a := activity.Activity{
				Name:        fmt.Sprintf("Green Activity %d", i),
				Description: "GreenCampus activity",
				Type:        types[i%len(types)],
				StartDate:   time.Now(),
				EndDate:     time.Now().AddDate(0, 1, 0),
				CreatedBy:   organizer.ID,
			}
			db.Create(&a)
		}
	}

	// ================= SUBMISSIONS =================
	db.Model(&activity.Submission{}).Count(&count)
	if count == 0 {
		var acts []activity.Activity
		var students []user.User

		db.Where("type = ?", "contest").Find(&acts)
		db.Where("role = ?", config.RoleStudent).Find(&students)

		for i := 0; i < 10 && i < len(acts) && i < len(students); i++ {
			s := activity.Submission{
				ActivityID: acts[i].ID,
				UserID:     students[i].ID,
				Content:    "This is a contest submission",
				Status:     "submitted",
				Score:      0,
			}
			db.Create(&s)
		}
	}

	// ================= SYSTEM CONFIG =================
	db.Model(&system.SystemConfig{}).Count(&count)
	if count == 0 {
		configs := []system.SystemConfig{
			{Key: "site_name", Value: "OU GreenCampus"},
			{Key: "email_sender", Value: "green@ou.edu.vn"},
			{Key: "max_event_capacity", Value: "500"},
			{Key: "contest_submission_limit", Value: "3"},
			{Key: "enable_registration", Value: "true"},
		}

		for _, cfg := range configs {
			db.Create(&cfg)
		}
	}

	// ================= AUDIT LOG =================
	db.Model(&system.AuditLog{}).Count(&count)
	if count == 0 {
		for i := 1; i <= 10; i++ {
			log := system.AuditLog{
				ActorID:   admin.ID,
				Action:    "SEED_DATA",
				Target:    fmt.Sprintf("INIT_%d", i),
				CreatedAt: time.Now(),
			}
			db.Create(&log)
		}
	}

	fmt.Println("✅ Seeding completed successfully")
	return nil
}
