package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	"net/http"
	"online-judge/common/errCode"
	"online-judge/models"
	"time"
)

type SubmissionRequest struct {
	PID      string `form:"pid" json:"pid" binding:"required"`
	UID      string `form:"uid" json:"uid" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	//Judge    string
	//Time     time.Time
	//Memory   int
	Language string `form:"language" json:"language" binding:"required"`
	//Created  time.Time
}

type Repdata struct {
	Data models.Problem `json:"data"`
}

type JudgeResponse struct {
	Stdout         string    `form:"stdout" json:"stdout" binding:"requird"`
	Time           time.Time `form:"time" json:"time" binding:"required"`
	Memory         string    `form:"memory" json:"memory" binding:"required"`
	Stderr         string    `form:"stderr" json:"stderr" binding:"required"`
	Token          string    `form:"token" json:"token" binding:"required"`
	Compile_output string    `form:"compile_output" json:"compile_output" binding:"required"`
	Message        string    `form:"message" json:"message" binding:"required"`
	//"status": {
	//"id": 3,
	//"description": "Accepted"
	//}
}

func PostSubmission(c *gin.Context) {
	req := SubmissionRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
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

	//Response(c, http.StatusOK, errCode.SUCCESS, data)

	res, err := grequests.Get("http://localhost:4040/api/v1/problem/detail", ro)

	data := Repdata{}
	res.JSON(&data)

	ro = &grequests.RequestOptions{
		Headers: map[string]string{
			"X-Auth-User":  "a1133bc6-a0f6-46bf-a2d8-6157418c6fe2",
			"X-Auth-Token": "f6583e60-b13b-4228-b554-2eb332ca64e7",
		},
		JSON: map[string]string{
			"source_code": req.Code,
			"language_id": req.Language,
			//"stdin": data.Data.SampleInput,
			//"expected_output": data.Data.SampleOutput,
			"stdin":           "Judge0",
			"expected_output": "hello, Judge0",
		},
	}

	res, err = grequests.Post("http://localhost:3000/submissions?wait=true", ro)
	if err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	Judgedata := JudgeResponse{}

	//Response(c, http.StatusOK, errCode.SUCCESS, data)

}
