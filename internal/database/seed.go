package database

import (
	"log"

	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/PNamGP1120/ougreencampus-go/pkg/hash"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	var count int64
	db.Model(&user.User{}).Count(&count)

	if count > 0 {
		log.Println("seed skipped: users already exist")
		return
	}

	password, _ := hash.HashPassword("admin123")

	admin := user.User{
		Email:    "admin@ou.edu.vn",
		Password: password,
		Name:     "System Admin",
		Role:     "admin",
		Status:   "active",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatal("failed to seed admin:", err)
	}

	log.Println("seed completed: admin user created")
}
