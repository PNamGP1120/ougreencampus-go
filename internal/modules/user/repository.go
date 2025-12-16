package user

import "gorm.io/gorm"

type Repository interface {
	Create(user *User) error
	FindAll() ([]User, error)
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Order("created_at desc").Find(&users).Error
	return users, err
}

func (r *repository) FindByID(id string) (*User, error) {
	var u User
	err := r.db.First(&u, "id = ?", id).Error
	return &u, err
}

func (r *repository) FindByEmail(email string) (*User, error) {
	var u User
	err := r.db.First(&u, "email = ?", email).Error
	return &u, err
}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}
