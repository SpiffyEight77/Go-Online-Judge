package controllers

import (
	"Go-Online-Judge/common/errCode"
	"Go-Online-Judge/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"github.com/spf13/viper"
	"net/http"
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
	Contest models.Contest          `json:"contest"`
	Problem []models.ContestProblem `json:"problem"`
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

	var list []models.ContestProblem

	for i := 0; i < len(pidList); i++ {
		data, _ := models.GetContestProblemDetail(pidList[i], contestID, i+1)
		list = append(list, *data)
	}

	problemDataResponse.Problem = list

	Response(c, http.StatusOK, errCode.SUCCESS, problemDataResponse)
}

type ContestCreateRequest struct {
	Title     string    `form:"title" json:"title" biding:"required"`
	StartTime time.Time `form:"start_time" json:"start_time"`
	EndTime   time.Time `form:"end_time" json:"end_time"`
	PIDList   string    `form:"pidList" json:"pid_list"`
	Problems  string    `form:"problem" json:"problems"`
	Type      string    `form:"type" json:"type"`
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
		fmt.Println(req)
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	contest := models.Contest{
		Title:     req.Title,
		Problems:  req.Problems,
		PIDList:   req.PIDList,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		//Type:      req.Type,
	}

	err, data := contest.ContestCreate()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	var pidList []int
	json.Unmarshal([]byte(req.PIDList), &pidList)

	//contestProblem := models.ContestProblem{}
	//problem := models.Problem{}

	cid := strconv.Itoa(data.ID)

	for i := 0; i < len(pidList); i++ {
		problem := models.Problem{
			ID: pidList[i],
		}

		//problem.ID = pidList[i]
		data, _ := problem.ProblemDetail()
		pid := strconv.Itoa(pidList[i])
		index := strconv.Itoa(i + 1)

		contestProblem := models.ContestProblem{
			CID:          cid,
			PID:          pid,
			Index:        index,
			Title:        data.Title,
			Description:  data.Description,
			Input:        data.Input,
			Output:       data.Output,
			SampleInput:  data.SampleInput,
			SampleOutput: data.SampleOutput,
		}
		_ = contestProblem.CreateContestProblem()
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

type ContestDeleteRequest struct {
	ID int `form:"id" json:"id"`
	//Title     string    `form:"title" json:"title" biding:"required"`
	//StartTime time.Time `form:"start_time" json:"start_time"`
	//EndTime   time.Time `form:"end_time" json:"end_time"`
	//PIDList   string    `form:"pidList" json:"pid_list"`
	//Problems  string    `form:"problem" json:"problems"`
	//Type      string    `form:"type" json:"type"`
}

// @Summary  Contest Detail
// @Produce json
// @Param contest_id query int true "contest_id"
// @Router /api/v1/admin/contest/delete [post]
func PostDeleteContest(c *gin.Context) {
	req := ContestDeleteRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	//contestID, err := strconv.Atoi(req.ID)
	//if err != nil {
	//	Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
	//	return
	//}

	contest := models.Contest{
		ID: req.ID,
	}
	if err := contest.ContestDelete(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

type ContestSubmissionRequest struct {
	PID      string `form:"pid" json:"pid" binding:"required"`
	UID      string `form:"uid" json:"uid" binding:"required"`
	Index    string `form:"index" json:"index" binding:"required"`
	CID      string `form:"cid" json:"cid" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	Language string `form:"language" json:"language" binding:"required"`
}

type RepContestData struct {
	Data models.Problem `json:"data"`
}

type ContestJudgeResponse struct {
	Stdout         string `form:"stdout" json:"stdout" binding:"required"`
	Time           string `form:"time" json:"time" binding:"required"`
	Memory         int    `form:"memory" json:"memory" binding:"required"`
	Stderr         string `form:"stderr" json:"stderr" binding:"required"`
	Token          string `form:"token" json:"token" binding:"required"`
	Compile_output string `form:"compile_output" json:"compile_output" binding:"required"`
	Message        string `form:"message" json:"message" binding:"required"`
	Status         struct {
		ID          int    `form:"id" json:"id" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
	} `form:"status" json:"status" binding:"required"`
}

func PostContestProblemSubmit(c *gin.Context) {
	req := ContestSubmissionRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	fmt.Println(req)

	programLanguage := map[string]string{
		"4":  "C (gcc 7.2.0)",
		"10": "C++ (g++ 7.2.0)",
		"28": "Java (OpenJDK 7)",
		"34": "Python (3.6.0)",
		"36": "Python (2.7.9)",
	}

	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"X-Auth-User":  "a1133bc6-a0f6-46bf-a2d8-6157418c6fe2",
			"X-Auth-Token": "f6583e60-b13b-4228-b554-2eb332ca64e7",
		},
		Params: map[string]string{
			//"problem_id": req.PID,
			"problem_index": req.Index,
			"contest_id":    req.CID,
		},
	}

	hostURL := viper.GetString("host.url")

	res, err := grequests.Get(hostURL+"/api/v1/contest/problem/detail", ro)
	data := RepContestData{}
	res.JSON(&data)

	fmt.Println(data)

	submitTime := time.Now()

	ro = &grequests.RequestOptions{
		Headers: map[string]string{
			"X-Auth-User":  "a1133bc6-a0f6-46bf-a2d8-6157418c6fe2",
			"X-Auth-Token": "f6583e60-b13b-4228-b554-2eb332ca64e7",
		},
		JSON: map[string]string{
			"source_code":     req.Code,
			"language_id":     req.Language,
			"stdin":           data.Data.SampleInput,
			"expected_output": data.Data.SampleOutput,
		},
	}

	judgeURL := viper.GetString("judge.url")
	judgeRes, err := grequests.Post(judgeURL+"/submissions?wait=true", ro)
	if err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}
	judgeData := ContestJudgeResponse{}
	judgeRes.JSON(&judgeData)

	contestSubmission := models.ContestSubmission{
		PID:       req.PID,
		UID:       req.UID,
		CID:       req.CID,
		Index:     req.Index,
		Code:      req.Code,
		Language:  programLanguage[req.Language],
		Username:  req.Username,
		Judge:     judgeData.Status.Description,
		Time:      judgeData.Time,
		Token:     judgeData.Token,
		Memory:    judgeData.Memory,
		CreatedAt: submitTime,
	}

	fmt.Println(contestSubmission)

	err = contestSubmission.CreateContestSubmission()
	if err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	//pid, err := strconv.Atoi(req.PID)
	contestProblem := models.ContestProblem{
		//Index:  req.Index,
		PID: contestSubmission.PID,
		CID: contestSubmission.CID,
	}

	if contestSubmission.Judge == "Accepted" {
		contestProblem.UpdateContestProblemSubmission(1, 1)
	} else {
		contestProblem.UpdateContestProblemSubmission(0, 1)
	}

	//uid, err := strconv.Atoi(req.UID)
	//user := models.User{
	//	ID: uid,
	//}

	//if contestSubmission.Judge == "Accepted" {
	//	user.UpdateUserSubmission(1, 1)
	//} else {
	//	user.UpdateUserSubmission(0, 1)
	//}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

func GetContestProblemDetail(c *gin.Context) {
	index := c.Query("problem_index")
	cid := c.Query("contest_id")

	if index == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	if cid == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	Index, _ := strconv.Atoi(index)
	CID, _ := strconv.Atoi(cid)

	contestProblem := models.ContestProblem{}

	data, err := contestProblem.GetContestProblemDetails(Index, CID)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

func GetContestSubmission(c *gin.Context) {
	cid := c.Query("cid")
	contestSubmission := models.ContestSubmission{
		CID: cid,
	}
	data, err := contestSubmission.ContestSubmissions()
	if err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

type UpdateContestRequest struct {
	ID        int       `form:"id" json:"id"`
	PIDList   string    `form:"pid_list" json:"pid_list"`
	Problems  string    `form:"problems" json:"problems"`
	StartTime time.Time `form:"start_time" json:"start_time"`
	EndTime   time.Time `form:"end_time" json:"end_time"`
}

func PostUpdateContest(c *gin.Context) {
	req := UpdateContestRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	contest := models.Contest{
		ID:        req.ID,
		PIDList:   req.PIDList,
		Problems:  req.Problems,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	if err := contest.ContestUpdate(); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
	return

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
