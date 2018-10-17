package models

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type Contest struct {
	ID          int       `gorm:"column:id" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	UID         int       `gorm:"column:uid" json:"uid"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	StartAt     time.Time `gorm:"column:start_at" json:"start_at"`
	EndAt       time.Time `gorm:"column:end_at" json:"end_at"`
	Status      int       `gorm:"column:status" json:"status"`
	ProblemNum  int       `gorm:"column:problem_num" json:"problem_num"`
	Participant int       `gorm:"column:participant" json:"participant"`
}

func (contest *Contest) ContestList() (*[]Contest, error) {
	var contestList []Contest

	key := "contestList"
	if Exists(key) {
		data, err := Get(key)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		json.Unmarshal(data, &contestList)
		return &contestList, nil
	}

	err := db.Model(&Contest{}).Scan(&contestList).Error
	if err != nil {
		return nil, err
	}
	Set(key, contestList, 3600)
	return &contestList, nil
}

func (contest *Contest) ContestDetail() (*Contest, error) {
	key := "contestID" + strconv.Itoa(contest.ID)
	if Exists(key) {
		data, err := Get(key)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		json.Unmarshal(data, &contest)
		return contest, nil
	}

	err := db.Model(&Contest{}).Where(&contest).Scan(&contest).Error
	if err != nil {
		return nil, err
	}
	Set(key, contest, 3600)
	return contest, nil
}

func (contest *Contest) ContestCreateAndUpdate() error {
	_, err := Delete("contestList")
	if err != nil {
		return err
	}

	if contest.ID == 0 {
		return db.Model(&Contest{}).Create(&contest).Error
	}

	key := "contestID" + strconv.Itoa(contest.ID)
	_, err = Delete(key)
	if err != nil {
		return err
	}
	return db.Model(&Contest{}).Update(&contest).Error
}

func (contest *Contest) ContestDelete() error {
	_, err := Delete("contestList")
	if err != nil {
		return err
	}

	key := "contestID" + strconv.Itoa(contest.ID)
	_, err = Delete(key)
	if err != nil {
		return err
	}
	return db.Model(&Contest{}).Delete(&contest).Error
}
