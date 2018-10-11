package models

type Solution struct {
	SID      int    `gorm:"column:sid" json:"sid"`
	PID      int    `gorm:"column:pid" json:"pid"`
	UID      int    `gorm:"column:uid" json:"uid"`
	Judge    int    `gorm:"column:judge" json:"judge"`
	Code     string `gorm:"column:code" json:"code"`
	Memory   int    `gorm:"column:memory" json:"memory"`
	Language int    `gorm:"column:language" json:"language"`
	Status   int    `gorm:"column:status" json:"status"`
}

func SubmitProblem(solution interface{}) error {
	return db.Model(&Solution{}).Create(&solution).Error
}

func SolutionList() (*[]Solution, error) {
	var solutionList []Solution
	return &solutionList, db.Model(&Solution{}).Scan(&solutionList).Error
}

func SolutionDetail(solutionID int) (*Solution, error) {
	var solution Solution
	return &solution, db.Model(&Solution{}).Where(&Solution{SID: solutionID}).Scan(&solution).Error
}
