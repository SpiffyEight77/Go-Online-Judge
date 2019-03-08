package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"github.com/spf13/viper"
	"net/http"
	"online-judge/common/errCode"
	"online-judge/models"
	"strconv"
	"time"
)

type SubmissionRequest struct {
	PID      string `form:"pid" json:"pid" binding:"required"`
	UID      string `form:"uid" json:"uid" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	Language string `form:"language" json:"language" binding:"required"`
}

type Repdata struct {
	Data models.Problem `json:"data"`
}

type JudgeResponse struct {
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

func PostSubmission(c *gin.Context) {
	req := SubmissionRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	fmt.Println(req.UID)

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
			"problem_id": req.PID,
		},
	}

	hostURL := viper.GetString("host.url")

	res, err := grequests.Get(hostURL+"/api/v1/problem/detail", ro)
	data := Repdata{}
	res.JSON(&data)

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
	judgeData := JudgeResponse{}
	judgeRes.JSON(&judgeData)

	submission := models.Submission{
		PID:       req.PID,
		UID:       req.UID,
		Code:      req.Code,
		Language:  programLanguage[req.Language],
		Username:  req.Username,
		Judge:     judgeData.Status.Description,
		Time:      judgeData.Time,
		Token:     judgeData.Token,
		Memory:    judgeData.Memory,
		CreatedAt: submitTime,
	}

	err = submission.CreateSubmission()
	if err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	pid, err := strconv.Atoi(req.PID)
	problem := models.Problem{
		ID: pid,
	}

	if submission.Judge == "Accepted" {
		problem.UpdateProblemSubmission(1, 1)
	} else {
		problem.UpdateProblemSubmission(0, 1)
	}

	uid, err := strconv.Atoi(req.UID)
	user := models.User{
		ID: uid,
	}

	if submission.Judge == "Accepted" {
		user.UpdateUserSubmission(1, 1)
	} else {
		user.UpdateUserSubmission(0, 1)
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

func GetSubmission(c *gin.Context) {
	submission := models.Submission{}
	data, err := submission.Submissions()
	if err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

type SolvedProblem struct {
	flag bool
}

func GetSolvedProblems(c *gin.Context) {
	pid := c.Query("pid")
	uid := c.Query("uid")
	if pid == "" || uid == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	submission := models.Submission{
		PID: pid,
		UID: pid,
	}

	data, _ := submission.SolvedSubmission()
	solvedProblem := SolvedProblem{}

	//if err != nil {
	//	Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
	//	return
	//}

	if data != nil {
		solvedProblem.flag = true
	} else {
		solvedProblem.flag = false
	}

	Response(c, http.StatusOK, errCode.SUCCESS, solvedProblem)
}
