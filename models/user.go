package models

import (
	"online-judge/controllers"
	"time"
)

type User struct {
	ID        int       `gorm:"column:id" json:"uid"`
	Username  string    `gorm:"column:username" json:"username"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	Token     string    `gorm:"column:token" json:"token"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	LastLogin time.Time `gorm:"column:last_login" json:"last_login"`
}

func CheckAuth(username, password string) (bool, *User) {
	var user User
	db.Model(&user).Where(User{Username: username, Password: password}).First(&user)
	if user.ID > 0 {
		return true, &user
	}
	return false, nil
}

func UpdateUserLogin(token string, lastLogin time.Time) error {
	var user User
	return db.Model(&user).Update(User{Token: token, LastLogin: lastLogin}).Error
}

func Register(username, password, email, token string) error {
	user := User{Username: username, Password: password, Email: email, Token: token, CreatedAt: time.Now(), LastLogin: time.Now()}
	return db.Model(&user).Create(&user).Error
}

func UserProfile(id int) (error, *User) {
	var user User
	return db.Model(&user).Where("id = ?", id).Scan(&user).Error, &user
}

func UpdateProfile(user controllers.UserProfileRequest) error {
	return db.Model(&user).Update().Error
}
