package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Faq(c *gin.Context) {
	c.HTML(http.StatusOK, "faq.html", gin.H{

	})
}
