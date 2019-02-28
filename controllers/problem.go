package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-judge/common/errCode"
	"online-judge/models"
	"strconv"
)

type ProblemRequest struct {
	Title        string `form:"title" json:"title" biding:"required"`
	Author       string `form:"author" json:"author" biding:"-"`
	Description  string `form:"descriptipon" json:"description" biding:"required"`
	Input        string `form:"input" json:"input" biding:"required"`
	Output       string `form:"output" json:"output" biding:"required"`
	SampleInput  string `form:"sample_input" json:"sample_input" biding:"required"`
	SampleOutput string `form:"sample_output" json:"sample_output" biding:"required"`
	Hint         string `form:"hint" json:"hint" biding:"-"`
}

// @Summary Problems List
// @Produce json
// @Router /api/v1/problem/list [get]
func GetProblems(c *gin.Context) {
	var problemList models.Problem
	data, err := problemList.ProblemsList()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

// @Summary Problem Detail
// @Produce json
// @Param problem_id query int true "problem_id"
// @Router /api/v1/problem/detail [get]
func GetProblemDetail(c *gin.Context) {
	pid := c.Query("problem_id")
	if pid == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	problemID, err := strconv.Atoi(pid)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	problem := models.Problem{
		ID: problemID,
	}
	data, err := problem.ProblemDetail()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

// @Summary Create Problem
// @Produce json
// @Param title query string true "title"
// @Param author query string true "author"
// @Param description query string true "description"
// @Param input query string true "input"
// @Param output query string true "output"
// @Param sample_input query string true "sample_input"
// @Param sample_output query string true "sample_output"
// @Param hint query string  false "hint"
// @Router /api/v1/problem/new [post]
//func PostCreateProblem(c *gin.Context) {
//	req := ProblemRequest{}
//	if err := c.ShouldBindJSON(&req); err != nil {
//		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
//		return
//	}
//
//	problem := models.Problem{
//		Title:        req.Title,
//		Author:       req.Author,
//		Description:  req.Description,
//		Input:        req.Input,
//		Output:       req.Output,
//		SampleInput:  req.SampleInput,
//		SampleOutput: req.SampleOutput,
//		Hint:         req.Hint,
//	}
//	if err := problem.CreateProblem(); err != nil {
//		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
//		return
//	}
//	Response(c, http.StatusOK, errCode.SUCCESS, nil)
//}

type ProblemDeleteRequest struct {
	IDList []int `form:"id_list" json:"id_list" biding:"required"`
}

// @Summary Delete Problem
// @Produce json
// @Param id_list query json true "id_list"
// @Router /api/v1/problem/delete [post]
func PostDeleteProblem(c *gin.Context) {
	req := ProblemDeleteRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	problem := models.Problem{
		IDList: req.IDList,
	}
	if err := problem.DeleteProblem(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

// @Summary Update Problem
// @Produce json
// @Param title query string true "title"
// @Param author query string true "author"
// @Param description query string true "description"
// @Param input query string true "input"
// @Param output query string true "output"
// @Param sample_input query string true "sample_input"
// @Param sample_output query string true "sample_output"
// @Param hint query string  false "hint"
// @Router /api/v1/problem/edit [post]
//func PostUpdateProblem(c *gin.Context) {
//	req := ProblemRequest{}
//	if err := c.ShouldBindJSON(&req); err != nil {
//		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
//		return
//	}
//
//	problem := models.Problem{
//		Title:        req.Title,
//		Author:       req.Author,
//		Description:  req.Description,
//		Input:        req.Input,
//		Output:       req.Output,
//		SampleInput:  req.SampleInput,
//		SampleOutput: req.SampleOutput,
//		Hint:         req.Hint,
//	}
//	if err := problem.UpdateProblem(); err != nil {
//		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
//		return
//	}
//	Response(c, http.StatusOK, errCode.SUCCESS, nil)
//}

type ProblemSubmitRequest struct {
	PID  int    `form:"pid" json:"pid" biding:"required"`
	UID  int    `form:"uid" json:"uid" biding:"required"`
	Code string `form:"code" json:"code" biding:"required"`
	//Memory   int    `form:"memory" json:"memory" biding:"required"`
	Language int `form:"language" json:"language" biding:"required"`
}

const (
	C       = 1
	Cpp     = 2
	Java    = 3
	Python2 = 4
	Python3 = 5
)

// @Summary Submit Problem
// @Produce json
// @Param pid query json true "pid"
// @Param uid query json true "uid"
// @Param code query json true "code"
// @Param memory query json true "memory"
// @Param language query json true "language"
// @Router /api/v1/problem/submit [post]
func PostSubmitProblem(c *gin.Context) {
	req := ProblemSubmitRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	solution := models.Solution{
		PID:  req.PID,
		UID:  req.UID,
		Code: req.Code,
		//Memory:   req.Memory,
		Language: req.Language,
	}

	if err := solution.SubmitProblem(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}
