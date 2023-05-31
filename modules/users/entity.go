package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Collections []Collection
}

type Collection struct {
	gorm.Model
	Name   string
	UserID uint
	User   User
}
