package models

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Problem struct {
	ID int `gorm:"column:id" json:"id"`
	//IDList       []int     `gorm:"column:id_list" json:"id_list"`
	Title        string `gorm:"column:title" json:"title"`
	Description  string `gorm:"column:description" json:"description"`
	Input        string `gorm:"column:input" json:"input"`
	Output       string `gorm:"column:output" json:"output"`
	SampleInput  string `gorm:"column:sample_input" json:"sample_input"`
	SampleOutput string `gorm:"column:sample_output" json:"sample_output"`
	//Hint         string    `gorm:"column:hint" json:"hint"`
	//Source       string    `gorm:"column:source" json:"source"`
	//Tags         string    `gorm:"column:tags" json:"tags"`
	//CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	//UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	Solve      int `gorm:"column:solve" json:"solve"`
	Submission int `gorm:"column:submission" json:"submission"`
}

func (problem *Problem) ProblemsList() (*[]Problem, error) {
	var problemList []Problem
	//key := "problemList"
	//if Exists(key) {
	//	data, err := Get(key)
	//	if err != nil {
	//		logs.Error(err)
	//		return nil, err
	//	}
	//	json.Unmarshal(data, &problemList)
	//	return &problemList, nil
	//}

	err := db.Model(&problemList).Scan(&problemList).Error
	if err != nil {
		return nil, err
	}
	//Set(key, problemList, 3600)
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

func GetProblemDetail(pid int) (*Problem, error) {
	var problem Problem
	err := db.Model(&Problem{}).Where(&problem).Scan(&problem).Error
	if err != nil {
		return nil, err
	}
	return &problem, nil
}

func (problem *Problem) CreateAndUpdateProblem() error {
	_, err := Delete("problemList")
	if err != nil {
		return err
	}

	if problem.ID == 0 {
		return db.Model(&Problem{}).Create(&problem).Error
	}

	key := "problemID" + strconv.Itoa(problem.ID)
	_, err = Delete(key)
	if err != nil {
		return err
	}
	return db.Model(&Problem{}).Update(&problem).Error
}

//func (problem *Problem) DeleteProblem() error {
//	_, err := Delete("problemList")
//	if err != nil {
//		return err
//	}
//
//	for k, _ := range problem.IDList {
//		key := "problemID" + strconv.Itoa(problem.IDList[k])
//		_, err = Delete(key)
//		if err != nil {
//			return err
//		}
//		if err := db.Model(&problem).Delete(&problem.IDList[k]).Error; err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (problem *Problem) UpdateProblemSubmission(solve, submission int) error {
	err := db.Model(&problem).UpdateColumn("submission", gorm.Expr("submission + ?", submission)).Error
	if err != nil {
		return err
	}

	err = db.Model(&problem).UpdateColumn("solve", gorm.Expr("solve + ?", solve)).Error
	if err != nil {
		return err
	}

	return nil
}

type ContestProblem struct {
	ID           int    `gorm:"column:id" json:"id"`
	CID          string `gorm:"column:cid" json:"cid"`
	PID          string `gorm:"column:pid" json:"pid"`
	Index        string `gorm:"column:index" json:"index"`
	Title        string `gorm:"column:title" json:"title"`
	Description  string `gorm:"column:description" json:"description"`
	Input        string `gorm:"column:input" json:"input"`
	Output       string `gorm:"column:output" json:"output"`
	SampleInput  string `gorm:"column:sample_input" json:"sample_input"`
	SampleOutput string `gorm:"column:sample_output" json:"sample_output"`
	Solve        int    `gorm:"column:solve" json:"solve"`
	Submission   int    `gorm:"column:submission" json:"submission"`
}

func (contestProblem *ContestProblem) UpdateContestProblemSubmission(solve, submission int) error {
	err := db.Model(&contestProblem).
		UpdateColumn("submission", gorm.Expr("submission + ?", submission)).
		Where("cid = ? and pid = ?", contestProblem.CID, contestProblem.ID).
		Error
	if err != nil {
		return err
	}

	err = db.Model(&contestProblem).
		UpdateColumn("solve", gorm.Expr("solve + ?", solve)).
		Where("cid = ? and pid = ?", contestProblem.CID, contestProblem.ID).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (contestProblem *ContestProblem) GetContestProblemDetails(index, cid int) (*ContestProblem, error) {
	//var contestProblem ContestProblem

	Index := strconv.Itoa(index)
	CID := strconv.Itoa(cid)
	//Index := strconv.Itoa(index)

	err := db.Model(&ContestProblem{}).
		Where(&ContestProblem{Index: Index, CID: CID}).
		Scan(&contestProblem).
		Error
	if err != nil {
		return nil, err
	}
	return contestProblem, nil
}

func GetContestProblemDetail(pid, cid, index int) (*ContestProblem, error) {
	var contestProblem ContestProblem

	PID := strconv.Itoa(pid)
	CID := strconv.Itoa(cid)
	Index := strconv.Itoa(index)

	err := db.Model(&ContestProblem{}).
		Where(&ContestProblem{PID: PID, CID: CID, Index: Index}).
		Scan(&contestProblem).
		Error
	if err != nil {
		return nil, err
	}
	return &contestProblem, nil
}

func (contestProblem *ContestProblem) CreateContestProblem() error {
	err := db.Model(&ContestProblem{}).Create(contestProblem).Error
	if err != nil {
		return err
	}
	return nil
}

func (problem *Problem) CreateProblem() error {
	err := db.Model(&Problem{}).Create(problem).Error
	if err != nil {
		return err
	}
	return nil
}

func (problem *Problem) UpdateProblem() error {
	err := db.Model(&Problem{}).Update(problem).Where("id = ?", problem.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (problem *Problem) DeleteProblem() error {
	err := db.Model(&Problem{}).Delete(problem).Where("id = ?", problem.ID).Error
	if err != nil {
		return err
	}
	return nil
}