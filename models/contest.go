package models

import (
	"strconv"
	"time"
)

type Contest struct {
	ID        int       `gorm:"column:id" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	PIDList   string    `gorm:"column:pidList" json:"pid_list"`
	Problems  string    `gorm:"column:problems" json:"problems"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	Type      string    `gorm:"column:type" json:"type"`
}

func (contest *Contest) ContestList() (*[]Contest, error) {
	var contestList []Contest
	err := db.Model(&Contest{}).Scan(&contestList).Error
	if err != nil {
		return nil, err
	}
	return &contestList, nil
}

func (contest *Contest) ContestDetail() (*Contest, error) {
	err := db.Model(&Contest{}).Where(&contest).Scan(&contest).Error
	if err != nil {
		return nil, err
	}
	return contest, nil
}

func (contest *Contest) ContestCreate() (error, *Contest) {
	err := db.Model(&contest).Create(&contest).Error
	if err != nil {
		return err, nil
	}
	return nil, contest
}

func (contest *Contest) ContestUpdate() error {
	err := db.Model(&contest).Update(&contest).Where("id = ?", contest.ID).Error
	if err != nil {
		return err
	}
	return nil
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
	//_, err := Delete("contestList")
	//if err != nil {
	//	return err
	//}

	//key := "contestID" + strconv.Itoa(contest.ID)
	//_, err = Delete(key)
	//if err != nil {
	//	return err
	//}
	return db.Model(&Contest{}).Delete(&contest).Where("id = ?",contest.ID).Error
}
