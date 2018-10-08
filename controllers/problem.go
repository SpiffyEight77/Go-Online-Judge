package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-judge/common/errCode"
	"online-judge/models"
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
	data, err := models.ProblemsList()
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
func PostCreateProblem(c *gin.Context) {
	req := ProblemRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	if err := models.CreateProblem(req); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

func PostDeleteProblem(c *gin.Context) {

}
