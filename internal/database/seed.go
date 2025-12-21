package database

import (
	"log"
	"time"

	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/PNamGP1120/ougreencampus-go/pkg/hash"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	log.Println("🌱 Seeding database...")

	seedUsers(db)
	seedSystemConfig(db)
	seedContent(db)
	seedActivity(db)
	seedEvent(db)

	log.Println("✅ Seed completed")
}

// ================= USERS =================

func seedUsers(db *gorm.DB) {
	password, _ := hash.HashPassword("123456")

	users := []user.User{
		{
			Name:     "Admin",
			Email:    "admin@ou.edu.vn",
			Password: password,
			Role:     "admin",
			Status:   "active",
		},
		{
			Name:     "Organizer",
			Email:    "organizer@ou.edu.vn",
			Password: password,
			Role:     "organizer",
			Status:   "active",
		},
		{
			Name:     "Student",
			Email:    "student@ou.edu.vn",
			Password: password,
			Role:     "student",
			Status:   "active",
		},
	}

	for _, u := range users {
		var count int64
		db.Model(&user.User{}).
			Where("email=?", u.Email).
			Count(&count)

		if count == 0 {
			db.Create(&u)
		}
	}
}

// ================= SYSTEM CONFIG =================

func seedSystemConfig(db *gorm.DB) {
	configs := []system.SystemConfig{
		{Key: "site_name", Value: "OU Green Campus"},
		{Key: "email_from", Value: "noreply@ou.edu.vn"},
		{Key: "event_checkin", Value: "qr"},
	}

	for _, c := range configs {
		var count int64
		db.Model(&system.SystemConfig{}).
			Where("key=?", c.Key).
			Count(&count)

		if count == 0 {
			db.Create(&c)
		}
	}
}

// ================= CONTENT =================

func seedContent(db *gorm.DB) {
	category := content.Category{Name: "Green Living"}
	db.FirstOrCreate(&category, content.Category{Name: "Green Living"})

	tag := content.Tag{Name: "Environment"}
	db.FirstOrCreate(&tag, content.Tag{Name: "Environment"})

	post := content.Content{
		Title:      "Welcome to OU Green Campus",
		Body:       "This is the first green article.",
		Image:      "",
		CategoryID: category.ID,
		AuthorID:   1,
		Tags:       []content.Tag{tag},
	}

	var count int64
	db.Model(&content.Content{}).
		Where("title=?", post.Title).
		Count(&count)

	if count == 0 {
		db.Create(&post)
	}
}

// ================= ACTIVITY =================

func seedActivity(db *gorm.DB) {
	act := activity.Activity{
		Title:       "Green Campaign 2025",
		Type:        "campaign",
		Description: "Planting trees around campus",
		Status:      "active",
		CreatedBy:   2,
	}

	var count int64
	db.Model(&activity.Activity{}).
		Where("title=?", act.Title).
		Count(&count)

	if count == 0 {
		db.Create(&act)
	}
}

// ================= EVENT =================

func seedEvent(db *gorm.DB) {
	ev := event.Event{
		Title:     "Green Workshop",
		Location:  "OU Main Hall",
		StartTime: time.Now().Add(48 * time.Hour),
		EndTime:   time.Now().Add(50 * time.Hour),
		CreatedBy: 2,
	}

	var count int64
	db.Model(&event.Event{}).
		Where("title=?", ev.Title).
		Count(&count)

	if count == 0 {
		db.Create(&ev)
	}
}
