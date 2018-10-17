package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-judge/common/errCode"
	"online-judge/models"
	"strconv"
	"time"
)

// @Summary  Contest List
// @Produce json
// @Router /api/v1/contest/list [get]
func GetContestList(c *gin.Context) {
	var contestList models.Contest
	data, err := contestList.ContestList()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

// @Summary  Contest Detail
// @Produce json
// @Param contest_id query int true "contest_id"
// @Router /api/v1/contest/detail [get]
func GetContestDetail(c *gin.Context) {
	cid := c.Query("contest_id")
	if cid == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	contestID, err := strconv.Atoi(cid)
	if err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	contest := models.Contest{
		ID: contestID,
	}
	data, err := contest.ContestDetail()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

type ContestCreateRequest struct {
	Title       string    `form:"title" json:"title" biding:"required"`
	UID         int       `form:"uid" json:"uid"  biding:"required"`
	StartAt     time.Time `form:"start_at" json:"start_at" biding:"required"`
	EndAt       time.Time `form:"end_at" json:"end_at" biding:"required"`
	Status      int       `form:"status" json:"status" biding:"required"`
	ProblemNum  int       `form:"problem_num" json:"problem_num" biding:"required"`
	Participant int       `form:"participant" json:"participant" biding:"required"`
}

// @Summary  Contest Create
// @Produce json
// @Param title query string true "title"
// @Param uid query int true "uid"
// @Param start_at query time true "start_at"
// @Param end_at query time true "end_at"
// @Param status query int true "status"
// @Param problem_num query int true "problem_num"
// @Param participant query int true "participant"
// @Router /api/v1/admin/contest/create [post]
func PostCreateContest(c *gin.Context) {
	req := ContestCreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	contest := models.Contest{
		Title:       req.Title,
		UID:         req.UID,
		StartAt:     req.StartAt,
		ProblemNum:  req.ProblemNum,
		Participant: req.Participant,
	}
	if err := contest.ContestCreateAndUpdate(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

// @Summary  Contest Detail
// @Produce json
// @Param contest_id query int true "contest_id"
// @Router /api/v1/admin/contest/delete [post]
func PostDeleteContest(c *gin.Context) {
	cid := c.Query("contest_id")
	if cid == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	contestID, err := strconv.Atoi(cid)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	contest := models.Contest{
		ID: contestID,
	}
	if err := contest.ContestDelete(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

// @Summary  Contest Edit
// @Produce json
// @Param title query string true "title"
// @Param uid query int true "uid"
// @Param start_at query time true "start_at"
// @Param end_at query time true "end_at"
// @Param status query int true "status"
// @Param problem_num query int true "problem_num"
// @Param participant query int true "participant"
// @Router /api/v1/admin/contest/edit [post]
func PostUpdateContest(c *gin.Context) {
	req := ContestCreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	contest := models.Contest{
		Title:       req.Title,
		UID:         req.UID,
		StartAt:     req.StartAt,
		ProblemNum:  req.ProblemNum,
		Participant: req.Participant,
	}
	if err := contest.ContestCreateAndUpdate(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}
