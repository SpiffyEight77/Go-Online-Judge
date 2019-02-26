package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	UserTableName     = "user"
	ProblemTableName  = "problem"
	ContestTableName  = "contest"
	SolutionTableName = "solution"
	NewsTableName     = "new"
)

var db *gorm.DB

func InitDB(dbType, dbURL string) {
	conn, err := gorm.Open(dbType, dbURL)
	if err != nil {
		panic(err)
	}

	db = conn

	db.SingularTable(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)
}
