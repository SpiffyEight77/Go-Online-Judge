package models

import "time"

type Submission struct {
	ID        int       `gorm:"column:id" json:"id"`
	UID       string    `gorm:"column:uid" json:"uid"`
	Username  string    `gorm:"column:username" json:"username"`
	PID       string    `gorm:"column:pid" json:"pid"`
	Judge     string    `gorm:"column:judge" json:"judge"`
	Code      string    `gorm:"column:code" json:"code"`
	Time      string    `gorm:"column:time" json:"time"`
	Memory    int       `gorm:"column:memory" json:"memory"`
	Language  string    `gorm:"column:language" json:"language"`
	Token     string    `gorm:"column:token" json:"token"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (submission *Submission) CreateSubmission() error {
	return db.Model(&Submission{}).Create(&submission).Error
}

func (submission *Submission) Submissions() (*[]Submission, error) {
	var submissions []Submission
	err := db.Model(&Submission{}).Scan(&submissions).Error
	if err != nil {
		return nil, err
	}
	return &submissions, nil
}
