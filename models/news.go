package models

import "time"

type News struct {
	ID        int       `form:"column:id" json:"id"`
	Content   string    `form:"column:content" json:"content"`
	CreatedAt time.Time `form:"column:created_at" json:"created_at"`
}

func NewsList() (*[]News, error) {
	var newsList []News
	return &newsList, db.Model(&News{}).Scan(&newsList).Error
}

func NewsDetail(newsID int) (News, error) {
	var news News
	return news, db.Model(&News{}).Where(&News{ID: newsID}).Scan(&news).Error
}

func NewsUpdate(news interface{}) error {
	return db.Model(&News{}).Update(&news).Error
}

func NewsCreate(news interface{}) error {
	return db.Model(&News{}).Create(&news).Error
}

func NewsDelete(newsID int) error {
	return db.Model(&News{}).Delete(&News{ID: newsID}).Error
}
