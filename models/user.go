package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type User struct {
	ID int `gorm:"column:id" json:"uid"`
	//IDList    []int     `gorm:"column:id_list" json:"id_list"`
	Nickname   string    `gorm:"column:nickname" json:"nickname"`
	Username   string    `gorm:"column:username" json:"username"`
	Email      string    `gorm:"column:email" json:"email"`
	Password   string    `gorm:"column:password" json:"password"`
	Submission int       `gorm:"column:submission" json:"submission"`
	Solve      int       `gorm:"column:solve" json:"solve"`
	Token      string    `gorm:"column:token" json:"token"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	LastLogin  time.Time `gorm:"column:last_login" json:"last_login"`
}

func (user *User) CheckAuth() (bool, *User) {
	var dbuser User
	db.Where("username = ?", user.Username).Find(&dbuser)
	if dbuser.ID > 0 {
		return true, &dbuser
	}
	return false, nil
}

func (user *User) Register() error {
	return db.Model(&User{}).Create(&user).Error
}

func (user *User) UpdateUserLogin() error {
	_, err := Delete("userList")
	if err != nil {
		return err
	}

	_, err = Delete("userID" + strconv.Itoa(user.ID))
	if err != nil {
		return nil
	}
	return db.Model(&user).Update(&user).Error
}

func (user *User) UserProfile() (*User, error) {
	//key := "userID" + strconv.Itoa(user.ID)
	//if Exists(key) {
	//	data, err := Get(key)
	//	if err != nil {
	//		logs.Error(err)
	//		return nil, err
	//	}
	//	json.Unmarshal(data, &user)
	//	return user, nil
	//}

	err := db.Model(&User{}).Where(&user).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	//Set(key, user, 3600)
	return user, nil
}

func (user *User) UpdateProfile() error {
	_, err := Delete("userList")
	if err != nil {
		return err
	}

	_, err = Delete("userID" + strconv.Itoa(user.ID))
	if err != nil {
		return nil
	}
	return db.Model(&User{}).Update(&user).Error
}

//func (user *User) DeleteUser() error {
//
//	_, err := Delete("userList")
//	if err != nil {
//		return err
//	}
//
//	for k, _ := range user.IDList {
//		_, err = Delete("userID" + strconv.Itoa(user.IDList[k]))
//		if err != nil {
//			return nil
//		}
//		if err := db.Model(&user).Delete(&user.IDList[k]).Error; err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (user *User) UserList() (*[]User, error) {
	var userList []User

	//key := "userList"
	//if Exists(key) {
	//	data, err := Get(key)
	//	if err != nil {
	//		logs.Error(err)
	//		return nil, err
	//	}
	//	json.Unmarshal(data, &userList)
	//	return &userList, nil
	//}

	err := db.Model(&User{}).Scan(&userList).Error
	if err != nil {
		return nil, err
	}
	//Set(key, userList, 3600)
	return &userList, nil
}

func (user *User) UpdateUserSubmission(solve, submission int) error {
	err := db.Model(&user).UpdateColumn("submission", gorm.Expr("submission + ?", submission)).Error
	if err != nil {
		return err
	}

	err = db.Model(&user).UpdateColumn("solve", gorm.Expr("solve + ?", solve)).Error
	if err != nil {
		return err
	}

	return nil
}
