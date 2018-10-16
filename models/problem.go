package models

import "time"

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
	return &problemList, db.Model(&problemList).Scan(&problemList).Error
}

func (problem *Problem) ProblemDetail() (*Problem, error) {
	return problem, db.Model(&Problem{}).Where(&problem).Scan(&problem).Error
}

func (problem *Problem) CreateProblem() error {
	return db.Model(&Problem{}).Create(&problem).Error
}

func (problem *Problem) UpdateProblem() error {
	return db.Model(&Problem{}).Update(&problem).Error
}

func DeleteProblem(idlist []int) error {
	problem := Problem{
		IDList: idlist,
	}
	for k, _ := range problem.IDList {
		if err := db.Model(&problem).Delete(&problem.IDList[k]).Error; err != nil {
			return err
		}
	}
	return nil
}
