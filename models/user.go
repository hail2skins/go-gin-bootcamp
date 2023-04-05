package models

import (
	"gin_notes/database"
	"gin_notes/helpers"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:64;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null;" json:""`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func CheckEmailAvailable(email string) bool {
	var user User
	database.Database.Where("username = ?", email).First(&user)
	return (user.ID == 0)
}

func UserCreate(email string, password string) *User {
	hshPasswd, _ := helpers.HashPassword(password)
	entry := User{Username: email, Password: hshPasswd}
	database.Database.Create(&entry)
	return &entry
}

func UserFind(id uint64) *User {
	var user User
	database.Database.Where("id = ?", id).First(&user)
	return &user
}

func UserFindByEmailAndPassword(email string, password string) *User {
	var user User
	database.Database.Where("username = ?", email).First(&user)
	if user.ID == 0 {
		return nil
	}
	match := helpers.CheckPasswordHash(password, user.Password)
	if match {
		return &user
	} else {
		return nil
	}
}
