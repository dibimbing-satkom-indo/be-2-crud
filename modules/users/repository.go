package users

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) Save(user *User) error {
	return r.db.Create(user).Error
}

func (r Repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
