package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-judge/common/errCode"
	"online-judge/models"
	"strconv"
)

// @Summary Solution List
// @Produce json
// @Router /api/v1/solution/list [get]
func GetSolutionList(c *gin.Context) {
	var solutionList models.Solution
	data, err := solutionList.SolutionList()
	if err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

// @Summary Solution List
// @Produce json
// @Param solution_id query int true "solution_id"
// @Router /api/v1/solution/detail [get]
func GetSolutionDetail(c *gin.Context) {
	sid := c.Query("solution_id")
	if sid == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	solutionID, err := strconv.Atoi(sid)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	solution := models.Solution{
		SID: solutionID,
	}
	data, err := solution.SolutionDetail()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}
