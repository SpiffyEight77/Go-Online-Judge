package models

import (
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

func (user *User) CheckAuth() (bool, *User) {
	db.Model(&user).Where(&user).First(&user)
	if user.ID > 0 {
		return true, user
	}
	return false, nil
}

func (user *User) Register() error {
	return db.Model(&User{}).Create(&user).Error
}

func (user *User) UpdateUserLogin() error {
	return db.Model(&user).Update(&user).Error
}

func (user *User) UserProfile() (*User, error) {
	return user, db.Model(&User{}).Where(&user).Scan(&user).Error
}

func (user *User) UpdateProfile() error {
	return db.Model(&User{}).Update(&user).Error
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

func (user *User) UserList() (*[]User, error) {
	var userList []User
	return &userList, db.Model(&User{}).Select("id,username,created_at,last_login").Scan(&userList).Error
}
