package main

import (
	"log"
	"os"
	"time"

	"github.com/PNamGP1120/ougreencampus-go/internal/database"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/PNamGP1120/ougreencampus-go/pkg/hash"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db := database.ConnectPostgres(dsn)
	log.Println("🌱 Seeding database...")

	// ================= USERS =================
	password, _ := hash.HashPassword("123456")

	users := []user.User{
		{Email: "admin@ou.edu.vn", Name: "Admin", Role: "admin", Password: password},
		{Email: "organizer@ou.edu.vn", Name: "Organizer", Role: "organizer", Password: password},
		{Email: "student@ou.edu.vn", Name: "Student", Role: "student", Password: password},
	}

	for _, u := range users {
		db.FirstOrCreate(&u, user.User{Email: u.Email})
	}

	var admin user.User
	db.First(&admin, "email = ?", "admin@ou.edu.vn")

	// ================= CONTENT =================
	categories := []content.Category{
		{Name: "News"},
		{Name: "Green Living"},
	}

	for _, c := range categories {
		db.FirstOrCreate(&c, content.Category{Name: c.Name})
	}

	tags := []content.Tag{
		{Name: "green"},
		{Name: "ou"},
	}

	for _, t := range tags {
		db.FirstOrCreate(&t, content.Tag{Name: t.Name})
	}

	contents := []content.Content{
		{
			Title:    "OU Green Campus Launch",
			Body:     "Green Campus officially launched",
			AuthorID: admin.ID,
		},
	}

	for _, c := range contents {
		db.FirstOrCreate(&c, content.Content{Title: c.Title})
	}

	// ================= ACTIVITY =================
	activities := []activity.Activity{
		{
			Title:       "Green Campaign 2025",
			Type:        "campaign",
			Description: "Reduce plastic usage",
		},
	}

	for _, a := range activities {
		db.FirstOrCreate(&a, activity.Activity{Title: a.Title})
	}

	// ================= EVENT =================
	events := []event.Event{
		{
			Title:       "Green Workshop",
			Description: "Workshop about sustainability",
			Location:    "OU Campus",
			StartTime:   time.Now().Add(24 * time.Hour),
			EndTime:     time.Now().Add(26 * time.Hour),
			Capacity:    100,
		},
	}

	for _, e := range events {
		db.FirstOrCreate(&e, event.Event{Title: e.Title})
	}

	// ================= SYSTEM CONFIG =================
	configs := []system.SystemConfig{
		{Key: "site_name", Value: "OU Green Campus"},
		{Key: "smtp_enabled", Value: "false"},
	}

	for _, c := range configs {
		db.FirstOrCreate(&c, system.SystemConfig{Key: c.Key})
	}

	log.Println("✅ Seed completed successfully")
}
