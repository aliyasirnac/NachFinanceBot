package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	BotId     int64
	FirstName string
	LastName  string
	Goals     []Goal
}

type Goal struct {
	gorm.Model
	UserID    int64
	Name      string
	Amount    float64
	StartDate time.Time
	EndDate   time.Time
}
