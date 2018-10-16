package models

import "time"

type News struct {
	ID        int       `form:"column:id" json:"id"`
	Content   string    `form:"column:content" json:"content"`
	CreatedAt time.Time `form:"column:created_at" json:"created_at"`
}

func (news *News) NewsList() (*[]News, error) {
	var newsList []News
	return &newsList, db.Model(&News{}).Scan(&newsList).Error
}

func (news *News) NewsDetail() (*News, error) {
	return news, db.Model(&News{}).Where(&news).Scan(&news).Error
}

func (news *News) NewsUpdate() error {
	return db.Model(&News{}).Update(&news).Error
}

func (news *News) NewsCreate() error {
	return db.Model(&News{}).Create(&news).Error
}

func (news *News) NewsDelete() error {
	return db.Model(&News{}).Delete(&news).Error
}
