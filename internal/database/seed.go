package database

import (
	"log"

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

	/* ---------- USERS ---------- */
	adminPass, _ := hash.HashPassword("admin123")
	orgPass, _ := hash.HashPassword("org123")
	stuPass, _ := hash.HashPassword("student123")

	admin := user.User{Email: "admin@ou.edu.vn", Name: "Admin", Role: "admin", Status: "active"}
	org := user.User{Email: "org@ou.edu.vn", Name: "Organizer", Role: "organizer", Status: "active"}
	stu := user.User{Email: "student@ou.edu.vn", Name: "Student", Role: "student", Status: "active"}

	db.FirstOrCreate(&admin, user.User{Email: admin.Email})
	db.FirstOrCreate(&org, user.User{Email: org.Email})
	db.FirstOrCreate(&stu, user.User{Email: stu.Email})

	db.Model(&admin).Update("password", adminPass)
	db.Model(&org).Update("password", orgPass)
	db.Model(&stu).Update("password", stuPass)

	/* ---------- CATEGORY ---------- */
	cat1 := content.Category{Name: "Green Life"}
	cat2 := content.Category{Name: "Environment"}

	db.FirstOrCreate(&cat1, content.Category{Name: cat1.Name})
	db.FirstOrCreate(&cat2, content.Category{Name: cat2.Name})

	/* ---------- TAG ---------- */
	tag1 := content.Tag{Name: "sustainability"}
	tag2 := content.Tag{Name: "climate"}

	db.FirstOrCreate(&tag1, content.Tag{Name: tag1.Name})
	db.FirstOrCreate(&tag2, content.Tag{Name: tag2.Name})

	/* ---------- CONTENT ---------- */
	post := content.Content{
		Title:      "Welcome to OU GreenCampus",
		Body:       "This is the first green article.",
		CategoryID: cat1.ID,
		AuthorID:   admin.ID,
		Tags:       []content.Tag{tag1, tag2},
	}
	db.FirstOrCreate(&post, content.Content{Title: post.Title})

	/* ---------- ACTIVITY ---------- */
	act := activity.Activity{
		Title:       "Green Campaign 2025",
		Type:        "campaign",
		Description: "Save the environment",
		OrganizerID: org.ID,
	}
	db.FirstOrCreate(&act, activity.Activity{Title: act.Title})

	/* ---------- EVENT ---------- */
	ev := event.Event{
		Title:       "Green Workshop",
		Time:        "2025-03-10 08:00",
		Location:    "OU Campus",
		OrganizerID: org.ID,
	}
	db.FirstOrCreate(&ev, event.Event{Title: ev.Title})

	/* ---------- SYSTEM CONFIG ---------- */
	cfgs := []system.SystemConfig{
		{Key: "site_name", Value: "OU GreenCampus"},
		{Key: "email_from", Value: "noreply@ou.edu.vn"},
	}

	for _, c := range cfgs {
		db.FirstOrCreate(&c, system.SystemConfig{Key: c.Key})
	}

	/* ---------- NOTIFICATION ---------- */
	noti := system.Notification{
		UserID: admin.ID,
		Title:  "Welcome",
		Body:   "Welcome to OU GreenCampus system",
	}
	db.FirstOrCreate(&noti, system.Notification{Title: noti.Title})

	log.Println("✅ Database seeded successfully")
}
