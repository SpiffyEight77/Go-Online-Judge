package controllers

import (
	"encoding/json"
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


type ProblemDataResponse struct {
	Contest models.Contest `json:"contest"`
	Problem []models.Problem `json:"problem"`
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

	var problemDataResponse ProblemDataResponse
	problemDataResponse.Contest = *data

	var pidList []int
	json.Unmarshal([]byte(data.PIDList), &pidList)
	//fmt.Println(pidList)

	var list []models.Problem

	for i := 0; i < len(pidList); i++ {
		data,_ := models.GetProblemDetail(pidList[i])
		//fmt.Println(data)
		list = append(list,*data)
	}

	problemDataResponse.Problem = list
	//fmt.Println(list)

	Response(c, http.StatusOK, errCode.SUCCESS, problemDataResponse)
}

type ContestCreateRequest struct {
	Title string `form:"title" json:"title" biding:"required"`
	//StartTime time.Time `form:"start_time" json:"start_time" biding:"required"`
	//StartTime time.Time `form:"start_time" json:"start_time"`
	//EndTime   time.Time `form:"end_time" json:"end_time" biding:"required"`
	//EndTime   time.Time `form:"end_time" json:"end_time"`
	PIDList string `form:"pidList" json:"pid_list" binding:"required"`
	Type    string `form:"type" json:"type" binding:"required"`
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
		Title:     req.Title,
		PIDList:   req.PIDList,
		StartTime: time.Now(),
		EndTime:   time.Now(),
		//StartTime: req.StartTime,
		//EndTime:   req.EndTime,
		Type: req.Type,
	}
	if err := contest.ContestCreate(); err != nil {
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
//func PostUpdateContest(c *gin.Context) {
//	req := ContestCreateRequest{}
//	if err := c.ShouldBindJSON(&req); err != nil {
//		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
//		return
//	}
//
//	contest := models.Contest{
//		Title:       req.Title,
//		UID:         req.UID,
//		StartAt:     req.StartAt,
//		ProblemNum:  req.ProblemNum,
//		Participant: req.Participant,
//	}
//	if err := contest.ContestCreateAndUpdate(); err != nil {
//		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
//		return
//	}
//	Response(c, http.StatusOK, errCode.SUCCESS, nil)
//}
