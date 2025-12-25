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

//
// ================= USERS =================
//

func seedUsers(db *gorm.DB) {
	password, _ := hash.HashPassword("123456")

	users := []user.User{
		{Name: "Admin OU", Email: "admin@ou.edu.vn", Password: password, Role: "admin", Status: "active"},
		{Name: "Green Organizer", Email: "organizer@ou.edu.vn", Password: password, Role: "organizer", Status: "active"},
		{Name: "Student A", Email: "student1@ou.edu.vn", Password: password, Role: "student", Status: "active"},
		{Name: "Student B", Email: "student2@ou.edu.vn", Password: password, Role: "student", Status: "active"},
		{Name: "Student C", Email: "student3@ou.edu.vn", Password: password, Role: "student", Status: "active"},
		{Name: "Student D", Email: "student4@ou.edu.vn", Password: password, Role: "student", Status: "active"},
		{Name: "Student E", Email: "student5@ou.edu.vn", Password: password, Role: "student", Status: "active"},
		{Name: "Volunteer 1", Email: "vol1@ou.edu.vn", Password: password, Role: "student", Status: "active"},
		{Name: "Volunteer 2", Email: "vol2@ou.edu.vn", Password: password, Role: "student", Status: "active"},
		{Name: "Volunteer 3", Email: "vol3@ou.edu.vn", Password: password, Role: "student", Status: "active"},
	}

	for _, u := range users {
		var count int64
		db.Model(&user.User{}).Where("email = ?", u.Email).Count(&count)
		if count == 0 {
			db.Create(&u)
		}
	}
}

//
// ================= SYSTEM CONFIG =================
//

func seedSystemConfig(db *gorm.DB) {
	configs := []system.SystemConfig{
		{Key: "site_name", Value: "OU Green Campus"},
		{Key: "email_from", Value: "noreply@ou.edu.vn"},
		{Key: "event_checkin", Value: "qr"},
		{Key: "theme", Value: "green"},
		{Key: "language", Value: "vi"},
		{Key: "maintenance", Value: "false"},
		{Key: "max_event_participants", Value: "500"},
		{Key: "allow_registration", Value: "true"},
		{Key: "footer_text", Value: "OU Green Campus – Sustainable Future"},
		{Key: "support_email", Value: "support@ou.edu.vn"},
	}

	for _, c := range configs {
		var count int64
		db.Model(&system.SystemConfig{}).Where("key = ?", c.Key).Count(&count)
		if count == 0 {
			db.Create(&c)
		}
	}
}

//
// ================= CONTENT =================
//

func seedContent(db *gorm.DB) {
	// Categories
	categories := []content.Category{
		{Name: "Green Living"},
		{Name: "Sustainability"},
		{Name: "Campus Life"},
	}

	for i := range categories {
		db.FirstOrCreate(&categories[i], content.Category{Name: categories[i].Name})
	}

	// Tags
	tags := []content.Tag{
		{Name: "Environment"},
		{Name: "Recycling"},
		{Name: "Energy"},
		{Name: "Water"},
		{Name: "Students"},
		{Name: "Campus"},
		{Name: "Climate"},
		{Name: "Lifestyle"},
		{Name: "Education"},
		{Name: "Community"},
	}

	for i := range tags {
		db.FirstOrCreate(&tags[i], content.Tag{Name: tags[i].Name})
	}

	// Contents (15+ bài, nội dung thật)
	posts := []content.Content{
		{
			Title:      "Welcome to OU Green Campus",
			Body:       "OU Green Campus is an initiative that promotes sustainability, environmental responsibility, and green living across the university community.",
			CategoryID: categories[0].ID,
			AuthorID:   1,
			Tags:       []content.Tag{tags[0], tags[9]},
		},
		{
			Title:      "Reducing Plastic Waste on Campus",
			Body:       "Plastic waste reduction is a key priority at OU. Students are encouraged to use reusable bottles, bags, and containers.",
			CategoryID: categories[0].ID,
			AuthorID:   2,
			Tags:       []content.Tag{tags[1], tags[4]},
		},
		{
			Title:      "Saving Energy in Classrooms",
			Body:       "Turning off lights, projectors, and air conditioners when not in use can significantly reduce energy consumption.",
			CategoryID: categories[1].ID,
			AuthorID:   1,
			Tags:       []content.Tag{tags[2], tags[6]},
		},
		{
			Title:      "Water Conservation Tips for Students",
			Body:       "Simple habits such as fixing leaks, using water efficiently, and reporting broken taps help conserve water resources.",
			CategoryID: categories[1].ID,
			AuthorID:   3,
			Tags:       []content.Tag{tags[3], tags[4]},
		},
		{
			Title:      "Green Transportation on Campus",
			Body:       "OU promotes walking, cycling, and public transportation as environmentally friendly commuting options.",
			CategoryID: categories[2].ID,
			AuthorID:   4,
			Tags:       []content.Tag{tags[5], tags[7]},
		},
		{
			Title:      "Tree Planting Activities",
			Body:       "Student volunteers regularly organize tree planting events to increase green spaces on campus.",
			CategoryID: categories[2].ID,
			AuthorID:   2,
			Tags:       []content.Tag{tags[0], tags[9]},
		},
		{
			Title:      "Recycling Guidelines at OU",
			Body:       "Understanding how to sort waste correctly is essential for the success of the campus recycling program.",
			CategoryID: categories[0].ID,
			AuthorID:   1,
			Tags:       []content.Tag{tags[1], tags[8]},
		},
		{
			Title:      "Sustainable Events Organization",
			Body:       "Event organizers are encouraged to minimize waste, avoid single-use plastics, and promote eco-friendly practices.",
			CategoryID: categories[1].ID,
			AuthorID:   2,
			Tags:       []content.Tag{tags[6], tags[9]},
		},
		{
			Title:      "Green Cafeteria Initiative",
			Body:       "The campus cafeteria is reducing food waste and promoting reusable food containers.",
			CategoryID: categories[2].ID,
			AuthorID:   3,
			Tags:       []content.Tag{tags[1], tags[7]},
		},
		{
			Title:      "Student Green Ambassadors",
			Body:       "Green Ambassadors play a key role in raising awareness about sustainability initiatives at OU.",
			CategoryID: categories[2].ID,
			AuthorID:   5,
			Tags:       []content.Tag{tags[4], tags[9]},
		},
		{
			Title:      "Climate Change Awareness",
			Body:       "Educational campaigns help students understand climate change and its impact on the environment.",
			CategoryID: categories[1].ID,
			AuthorID:   1,
			Tags:       []content.Tag{tags[6], tags[8]},
		},
		{
			Title:      "Smart Energy Monitoring Systems",
			Body:       "OU has implemented smart monitoring systems to track and optimize energy usage across buildings.",
			CategoryID: categories[1].ID,
			AuthorID:   2,
			Tags:       []content.Tag{tags[2], tags[6]},
		},
		{
			Title:      "Community Clean-up Campaign",
			Body:       "Students and staff participate in clean-up campaigns to keep the campus clean and green.",
			CategoryID: categories[0].ID,
			AuthorID:   4,
			Tags:       []content.Tag{tags[0], tags[9]},
		},
		{
			Title:      "Green Lifestyle for Students",
			Body:       "Adopting a green lifestyle helps students reduce their environmental footprint and live sustainably.",
			CategoryID: categories[0].ID,
			AuthorID:   5,
			Tags:       []content.Tag{tags[7], tags[4]},
		},
		{
			Title:      "Sustainable Future at OU",
			Body:       "OU Green Campus aims to create a sustainable future through education, innovation, and community engagement.",
			CategoryID: categories[1].ID,
			AuthorID:   1,
			Tags:       []content.Tag{tags[6], tags[9]},
		},
	}

	for _, post := range posts {
		var count int64
		db.Model(&content.Content{}).Where("title = ?", post.Title).Count(&count)
		if count == 0 {
			db.Create(&post)
		}
	}
}

//
// ================= ACTIVITY =================
//

func seedActivity(db *gorm.DB) {
	activities := []activity.Activity{
		{Title: "Tree Planting Campaign", Type: "campaign", Description: "Planting trees around campus", Status: "active", CreatedBy: 2},
		{Title: "Plastic Free Week", Type: "campaign", Description: "Reducing plastic usage", Status: "active", CreatedBy: 2},
		{Title: "Energy Saving Challenge", Type: "challenge", Description: "Saving electricity in dormitories", Status: "active", CreatedBy: 1},
		{Title: "Water Conservation Drive", Type: "campaign", Description: "Promoting water saving habits", Status: "active", CreatedBy: 3},
		{Title: "Green Ambassador Training", Type: "training", Description: "Training student ambassadors", Status: "active", CreatedBy: 1},
		{Title: "Campus Clean-up Day", Type: "event", Description: "Cleaning campus areas", Status: "active", CreatedBy: 2},
		{Title: "Recycling Workshop", Type: "workshop", Description: "Learning recycling techniques", Status: "active", CreatedBy: 1},
		{Title: "Bike to Campus Day", Type: "campaign", Description: "Encouraging cycling", Status: "active", CreatedBy: 3},
		{Title: "Eco Innovation Contest", Type: "contest", Description: "Innovative green ideas", Status: "active", CreatedBy: 1},
		{Title: "Green Volunteer Meetup", Type: "meetup", Description: "Volunteer networking", Status: "active", CreatedBy: 2},
	}

	for _, act := range activities {
		var count int64
		db.Model(&activity.Activity{}).Where("title = ?", act.Title).Count(&count)
		if count == 0 {
			db.Create(&act)
		}
	}
}

//
// ================= EVENT =================
//

func seedEvent(db *gorm.DB) {
	events := []event.Event{
		{Title: "Green Campus Workshop", Location: "Main Hall", StartTime: time.Now().Add(24 * time.Hour), EndTime: time.Now().Add(26 * time.Hour), CreatedBy: 2},
		{Title: "Tree Planting Day", Location: "Campus Garden", StartTime: time.Now().Add(48 * time.Hour), EndTime: time.Now().Add(50 * time.Hour), CreatedBy: 2},
		{Title: "Sustainability Seminar", Location: "Room A101", StartTime: time.Now().Add(72 * time.Hour), EndTime: time.Now().Add(74 * time.Hour), CreatedBy: 1},
		{Title: "Recycling Training", Location: "Lab 3", StartTime: time.Now().Add(96 * time.Hour), EndTime: time.Now().Add(98 * time.Hour), CreatedBy: 1},
		{Title: "Green Innovation Fair", Location: "Outdoor Stage", StartTime: time.Now().Add(120 * time.Hour), EndTime: time.Now().Add(124 * time.Hour), CreatedBy: 2},
		{Title: "Water Saving Workshop", Location: "Room B202", StartTime: time.Now().Add(144 * time.Hour), EndTime: time.Now().Add(146 * time.Hour), CreatedBy: 3},
		{Title: "Eco Lifestyle Talk", Location: "Auditorium", StartTime: time.Now().Add(168 * time.Hour), EndTime: time.Now().Add(170 * time.Hour), CreatedBy: 1},
		{Title: "Campus Clean-up Event", Location: "Main Gate", StartTime: time.Now().Add(192 * time.Hour), EndTime: time.Now().Add(194 * time.Hour), CreatedBy: 2},
		{Title: "Green Volunteer Orientation", Location: "Meeting Room", StartTime: time.Now().Add(216 * time.Hour), EndTime: time.Now().Add(218 * time.Hour), CreatedBy: 3},
		{Title: "Sustainable Future Forum", Location: "Conference Hall", StartTime: time.Now().Add(240 * time.Hour), EndTime: time.Now().Add(242 * time.Hour), CreatedBy: 1},
	}

	for _, ev := range events {
		var count int64
		db.Model(&event.Event{}).Where("title = ?", ev.Title).Count(&count)
		if count == 0 {
			db.Create(&ev)
		}
	}
}
