package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"online-judge/common/errCode"
	"online-judge/models"
	"strconv"
	"time"
)

// @Summary News List
// @Produce json
// @Router /api/v1/news/list [get]
func GetNewsList(c *gin.Context) {
	var newsList models.News
	data, err := newsList.NewsList()
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

	news := models.News{
		ID: newsID,
	}
	data, err := news.NewsDetail()
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

	news := models.News{
		Content: req.Content,
	}
	if err := news.NewsCreateAndUpdate(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

// @Summary  News Create
// @Produce json
// @Param content query string true "content"
// @Router /api/v1/news/create [post]
//func PostNewsCreate(c *gin.Context) {
//	req := NewsEditRequest{}
//	if err := c.ShouldBindJSON(&req); err != nil {
//		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
//		return
//	}
//
//	news := models.News{
//		Content: req.Content,
//	}
//	if err := news.NewsCreateAndUpdate(); err != nil {
//		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
//		return
//	}
//	Response(c, http.StatusOK, errCode.SUCCESS, nil)
//}

type NewsRequest struct {
	NID       int       `form:"nid" json:"nid"`
	Title     string    `form:"title" json:"title"`
	Content   string    `form:"content" json:"content"`
	CreatedAt time.Time `form:"created_at" json:"created_at"`
}

func PostNewsUpdate(c *gin.Context) {
	req := NewsRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	fmt.Println(req)

	news := models.News{
		ID:      req.NID,
		Title:   req.Title,
		Content: req.Content,
		//CreatedAt: req.CreatedAt,
	}

	if err := news.NewsUpdate(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

func PostNewsCreate(c *gin.Context) {
	req := NewsRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	fmt.Println(req)

	news := models.News{
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: req.CreatedAt,
	}

	if err := news.NewsCreate(); err != nil {
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
	req := NewsRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}
	news := models.News{
		ID: req.NID,
	}
	if err := news.NewsDelete(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}
