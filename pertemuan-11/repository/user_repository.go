package repository

import (
	"pertemuan-11/config"
	"pertemuan-11/model"
)

func GetAllUser() ([]model.User, error) {
	var users []model.User
	
	// GORM: SELECT * FROM users
	err := config.DB.Find(&users).Error
	
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindByUsername(username string) (model.User, error) {
	var user model.User
	
	// GORM: SELECT * FROM users WHERE username = '...' LIMIT 1
	err := config.DB.Where("username = ?", username).First(&user).Error

	return user, err
}

func CreateUser(user model.User) (model.User, error) {
	// GORM: INSERT INTO users (...) VALUES (...)
	err := config.DB.Create(&user).Error
	
	if err != nil {
		return user, err
	}

	return user, nil
}