package user

import "gorm.io/gorm"

type Repository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) error
	List() ([]User, error)
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

func (r *repository) FindByEmail(email string) (*User, error) {
	var u User
	err := r.db.Where("email = ?", email).First(&u).Error
	return &u, err
}

func (r *repository) FindByID(id uint) (*User, error) {
	var u User
	err := r.db.First(&u, id).Error
	return &u, err
}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *repository) List() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}
