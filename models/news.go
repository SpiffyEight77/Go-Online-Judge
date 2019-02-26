package models

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type News struct {
	ID        int       `form:"column:id" json:"id"`
	Title     string    `form:"column:title" json:"title"`
	Content   string    `form:"column:content" json:"content"`
	CreatedAt time.Time `form:"column:created_at" json:"created_at"`
}

func (news *News) NewsList() (*[]News, error) {
	var newsList []News
	//key := "newsList"
	//if Exists(key) {
	//	data, err := Get(key)
	//	if err != nil {
	//		logs.Error(err)
	//		return nil, err
	//	}
	//	json.Unmarshal(data, &newsList)
	//}

	err := db.Model(&News{}).Scan(&newsList).Error
	if err != nil {
		return nil, err
	}
	//Set(key, newsList, 3600)
	return &newsList, nil
}

func (news *News) NewsDetail() (*News, error) {
	key := "newsID" + strconv.Itoa(news.ID)
	if Exists(key) {
		data, err := Get(key)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		json.Unmarshal(data, &news)
		return news, nil
	}

	err := db.Model(&News{}).Where(&news).Scan(&news).Error
	if err != nil {
		return nil, err
	}
	Set(key, news, 3600)
	return news, nil
}

func (news *News) NewsCreateAndUpdate() error {
	_, err := Delete("newsList")
	if err != nil {
		return err
	}

	if news.ID == 0 {
		return db.Model(&News{}).Create(&news).Error
	}

	key := "newsID" + strconv.Itoa(news.ID)
	_, err = Delete(key)
	if err != nil {
		return err
	}
	return db.Model(&News{}).Update(&news).Error
}

func (news *News) NewsDelete() error {
	_, err := Delete("newsList")
	if err != nil {
		return err
	}

	key := "newsID" + strconv.Itoa(news.ID)
	_, err = Delete(key)
	if err != nil {
		return err
	}
	return db.Model(&News{}).Delete(&news).Error
}
