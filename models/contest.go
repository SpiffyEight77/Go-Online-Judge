package models

import "time"

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

func ContestList() (*[]Contest, error) {
	var contestList []Contest
	return &contestList, db.Model(&Contest{}).Scan(&contestList).Error
}

func ContestDetail(contestID int) (*Contest, error) {
	var contest Contest
	return &contest, db.Model(&Contest{}).Where(&Contest{ID: contestID}).Scan(&contest).Error
}

func ContestCreate(contest interface{}) error {
	return db.Model(&Contest{}).Create(&contest).Error
}

func ContestDelete(contestID int) error {
	return db.Model(&Contest{}).Delete(&Contest{ID: contestID}).Error
}

func ContestUpdate(contest interface{}) error {
	return db.Model(&Contest{}).Update(&contest).Error
}
