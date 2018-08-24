package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	"fmt"
)

type UserRegisterRequest struct {
	Username 		 string `form:"username"        json:"username"        binding:"required"`
	Password 		 string `form:"password"        json:"password"        binding:"required"`
	Confirmpassword  string `form:"confirmpassword" json:"confirmpassword" binding:"required"`
	//Nickname 		 string `form:"nickname"        json:"nickname"        binding:"required"`
	Email 		     string `form:"email"           json:"email"           binding:"required"`
}

//@GET
func Login (c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{

	})
}
//@GET
func Register (c *gin.Context) {
 	c.HTML(http.StatusOK, "register.html", gin.H{

	})
}

//@POST
func PostRegister (c *gin.Context) {
	req := UserRegisterRequest{}
	if err := c.MustBindWith(&req, binding.Form); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(req)
		c.HTML(http.StatusOK, "OK.html", gin.H{

		})
	}

}
