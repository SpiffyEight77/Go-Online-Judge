package models

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type Problem struct {
	ID           int       `gorm:"column:id" json:"id"`
	IDList       []int     `gorm:"column:id_list" json:"id_list"`
	Title        string    `gorm:"column:title" json:"title"`
	Author       string    `gorm:"column:author" json:"author"`
	Description  string    `gorm:"column:description" json:"description"`
	Input        string    `gorm:"column:input" json:"input"`
	Output       string    `gorm:"column:output" json:"output"`
	SampleInput  string    `gorm:"column:sample_input" json:"sample_input"`
	SampleOutput string    `gorm:"column:sample_output" json:"sample_output"`
	Hint         string    `gorm:"column:hint" json:"hint"`
	Source       string    `gorm:"column:source" json:"source"`
	Tags         string    `gorm:"column:tags" json:"tags"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (problem *Problem) ProblemsList() (*[]Problem, error) {
	var problemList []Problem
	key := "problemList"
	if Exists(key) {
		data, err := Get(key)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		json.Unmarshal(data, &problemList)
		return &problemList, nil
	}

	err := db.Model(&problemList).Scan(&problemList).Error
	if err != nil {
		return nil, err
	}
	Set(key, problemList, 3600)
	return &problemList, nil
}

func (problem *Problem) ProblemDetail() (*Problem, error) {
	key := "problemID" + strconv.Itoa(problem.ID)
	if Exists(key) {
		data, err := Get(key)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		json.Unmarshal(data, &problem)
		return problem, nil
	}

	err := db.Model(&Problem{}).Where(&problem).Scan(&problem).Error
	if err != nil {
		return nil, err
	}
	Set(key, problem, 3600)
	return problem, nil
}

func (problem *Problem) CreateProblem() error {
	_, err := Delete("problemList")
	if err != nil {
		return err
	}

	return db.Model(&Problem{}).Create(&problem).Error
}

func (problem *Problem) UpdateProblem() error {
	_, err := Delete("problemList")
	if err != nil {
		return err
	}

	key := "problemID" + strconv.Itoa(problem.ID)
	_, err = Delete(key)
	if err != nil {
		return err
	}
	return db.Model(&Problem{}).Update(&problem).Error
}

func (problem *Problem) DeleteProblem() error {
	_, err := Delete("problemList")
	if err != nil {
		return err
	}

	for k, _ := range problem.IDList {
		key := "problemID" + strconv.Itoa(problem.IDList[k])
		_, err = Delete(key)
		if err != nil {
			return err
		}
		if err := db.Model(&problem).Delete(&problem.IDList[k]).Error; err != nil {
			return err
		}
	}
	return nil
}
