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

func (solution *Solution) SubmitProblem() error {
	return db.Model(&Solution{}).Create(&solution).Error
}

func (solution *Solution) SolutionList() (*[]Solution, error) {
	var solutionList []Solution
	return &solutionList, db.Model(&Solution{}).Scan(&solutionList).Error
}

func (solution *Solution) SolutionDetail() (*Solution, error) {
	return solution, db.Model(&Solution{}).Where(&solution).Scan(&solution).Error
}
