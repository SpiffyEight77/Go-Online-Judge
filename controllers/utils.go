package controllers

import (
	"Go-Online-Judge/common/errCode"
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpCode, code int, data interface{}) {
	if httpCode == 500 {
		c.Status(500)
	}

	c.JSON(httpCode, gin.H{
		"msg":  errCode.ErrorString(code),
		"code": code,
		"data": data,
	})
}
