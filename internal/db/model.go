package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	BotId       int64
	FirstName   string
	LastName    string
	PhoneNumber string
}

