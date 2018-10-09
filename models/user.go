package models

import (
	"fmt"
	"time"
)

type User struct {
	ID        int       `gorm:"column:id" json:"uid"`
	IDList    []int     `gorm:"column:id_list" json:"id_list"`
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
	user := User{
		Username: username,
		Password: password,
		Email: email,
		Token: token,
		CreatedAt: time.Now(),
		LastLogin: time.Now(),
	}
	return db.Model(&user).Create(&user).Error
}

func UserProfile(id int) (error, *User) {
	var user User
	return db.Model(&user).Where("id = ?", id).Scan(&user).Error, &user
}

func UpdateProfile(user interface{}) error {
	fmt.Println(user)
	return db.Model(&user).Update().Error
}

func DeleteUser(idlist []int) error {
	user := User{
		IDList: idlist,
	}
	for k, _ := range user.IDList {
		if err := db.Model(&user).Delete(&user.IDList[k]).Error; err != nil {
			return err
		}
	}
	return nil
}
