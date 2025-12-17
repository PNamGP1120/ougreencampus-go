package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var u User
	return &u, r.db.Where("email = ?", email).First(&u).Error
}

func (r *Repository) FindByID(id string) (*User, error) {
	var u User
	return &u, r.db.First(&u, "id = ?", id).Error
}

func (r *Repository) List() ([]User, error) {
	var users []User
	return users, r.db.Find(&users).Error
}

func (r *Repository) Create(u *User) error {
	return r.db.Create(u).Error
}

func (r *Repository) Update(u *User) error {
	return r.db.Save(u).Error
}
