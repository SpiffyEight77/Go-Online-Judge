package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	UserTableName    = "user"
	ProblemTableName = "problem"
)

func InitDB(dbType, dbURL string) {

	db, err := gorm.Open(dbType, dbURL)
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)

}
