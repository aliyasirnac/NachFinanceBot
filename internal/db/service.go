package db

import (
	"gorm.io/gorm"
)

type Service interface {
	GetUser(id int64) (User, error)
	GetGoals(id int64) ([]Goal, error)
	AddGoal(goal Goal) error
	UpdateGoal(goal Goal) error
	DeleteGoal(id int64) error
	AddUser(user User) error
	UpdateUser(user User) error
	DeleteUser(id int64) error
	GetUsers() ([]User, error)
	GetUserGoals(id int64) ([]Goal, error)
	GetUserGoal(userID int64, goalID int64) (Goal, error)
	GetUserByBotId(id int64) (User, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return service{db: db}
}

func (s service) GetUserByBotId(id int64) (User, error) {
	var user User
	result := s.db.Where("bot_id = ?", id).First(&user)
	return user, result.Error
}

func (s service) GetUser(id int64) (User, error) {
	var user User
	result := s.db.First(&user, id)
	return user, result.Error
}

func (s service) GetGoals(id int64) ([]Goal, error) {
	var goals []Goal
	result := s.db.Where("user_id = ?", id).Find(&goals)
	return goals, result.Error
}

func (s service) AddGoal(goal Goal) error {
	result := s.db.Create(&goal)
	return result.Error
}

func (s service) UpdateGoal(goal Goal) error {
	result := s.db.Save(&goal)
	return result.Error
}

func (s service) DeleteGoal(id int64) error {
	result := s.db.Delete(&Goal{}, id)
	return result.Error
}

func (s service) AddUser(user User) error {
	result := s.db.Create(&user)
	return result.Error
}

func (s service) UpdateUser(user User) error {
	result := s.db.Save(&user)
	return result.Error
}

func (s service) DeleteUser(id int64) error {
	result := s.db.Delete(&User{}, id)
	return result.Error
}

func (s service) GetUsers() ([]User, error) {
	var users []User
	result := s.db.Find(&users)
	return users, result.Error
}

func (s service) GetUserGoals(id int64) ([]Goal, error) {
	return s.GetGoals(id)
}

func (s service) GetUserGoal(userID int64, goalID int64) (Goal, error) {
	var goal Goal
	result := s.db.Where("user_id = ? AND id = ?", userID, goalID).First(&goal)
	return goal, result.Error
}
