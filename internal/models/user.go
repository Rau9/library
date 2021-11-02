package models

import (
	"gorm.io/gorm"
)

type User struct {
	Base
	Username string `gorm:"size:50;not null;index;unique"`
	Password string `gorm:"size:250;not null"`
}

func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

func CreateUserIfNotExists(dbclient *gorm.DB, username string, password string) (*User, error) {
	user := NewUser(username, password)
	result := dbclient.FirstOrCreate(user, User{Username: username, Password: password})
	return user, result.Error
}

func GetUserByName(dbclient *gorm.DB, username string) (*User, error) {
	user := &User{}
	result := dbclient.Where("username = ?", username).First(user)
	return user, result.Error
}
