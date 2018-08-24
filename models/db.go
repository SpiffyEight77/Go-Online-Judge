package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	UserInfoTableName = "user_info"
)

func Initdb() {

	db, err := gorm.Open("mysql", "root:root@/Eu_OJ?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	if !db.HasTable(&UserInfo{}) {
		db.AutoMigrate(&UserInfo{})
	}

	defer db.Close()

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)

}
