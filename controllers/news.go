package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-judge/common/errCode"
	"online-judge/models"
	"strconv"
)

// @Summary News List
// @Produce json
// @Router /api/v1/news/list [get]
func GetNewsList(c *gin.Context) {
	data, err := models.NewsList()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

// @Summary  News Detail
// @Produce json
// @Param news_id query int true "news_id"
// @Router /api/v1/news/detail [get]
func GetNewsDetail(c *gin.Context) {
	nid := c.Query("news_id")
	if nid == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	newsID, err := strconv.Atoi(nid)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	data, err := models.NewsDetail(newsID)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

type NewsEditRequest struct {
	Content string `form:"column:content" json:"content" biding:"required"`
}

// @Summary  News Edit
// @Produce json
// @Param content query string true "content"
// @Router /api/v1/news/edit [post]
func PostNewsEdit(c *gin.Context) {
	req := NewsEditRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	if err := models.NewsUpdate(req); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

// @Summary  News Create
// @Produce json
// @Param content query string true "content"
// @Router /api/v1/news/create [post]
func PostNewsCreate(c *gin.Context) {
	req := NewsEditRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	if err := models.NewsCreate(req); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

// @Summary  News Delete
// @Produce json
// @Param news_id query int true "news_id"
// @Router /api/v1/news/delete [post]
func PostNewsDelete(c *gin.Context) {
	nid := c.Query("news_id")
	if nid == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	newsID, err := strconv.Atoi(nid)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	if err := models.NewsDelete(newsID); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}
