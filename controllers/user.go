package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"online-judge/common/errCode"
	"online-judge/middlewares/jwt"
	"online-judge/models"
	"time"
)

type UserLoginRequest struct {
	Username string `form:"username" json:"username" biding:"required"`
	Password string `form:"password" json:"password" biding:"required"`
}

// @Summary User & Admin Login
// @Produce json
// @Router /api/v1/admin/user/login [get]
func GetUserLogin(c *gin.Context) {
	Response(c, 200, errCode.SUCCESS, nil)
}

// @Summary User & Admin Login
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Router /api/v1/admin/user/login [post]
func PostUserLogin(c *gin.Context) {
	var req UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	valid := validation.Validation{}
	ok, _ := valid.Valid(&req)

	data := make(map[string]interface{})
	code := errCode.BADREQUEST

	if ok {
		if isExist, data := models.CheckAuth(req.Username, req.Password); isExist == true {
			if token, err := jwt.GenerateToken(req.Username, req.Password); err != nil {
				code = errCode.UNAUTHORIZED
			} else {
				data.Token = token
				code = errCode.SUCCESS
				models.UpdateUserLogin(token, time.Now())
			}
		} else {
			code = errCode.UNAUTHORIZED
		}
	}
	Response(c, code, code, data)
}

type UserRegisterRequest struct {
	Username string `form:"username" json:"username" biding:"required"`
	Email    string `form:"email" json:"email" biding:"required"`
	Password string `form:"password" json:"password" biding:"required"`
}

// @Summary User Register
// @Produce json
// @Router /api/v1/user/login [get]
func GetUserRegister(c *gin.Context) {
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

// @Summary User Register
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Router /api/v1/user/login [post]
func PostUserRegister(c *gin.Context) {
	req := UserRegisterRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	if ok, data := models.CheckAuth(req.Username, req.Password); ok || data != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	token, _ := jwt.GenerateToken(req.Username, req.Password)
	if err := models.Register(req.Username, req.Password, req.Email, token); err == nil {
		if ok, data := models.CheckAuth(req.Username, req.Password); ok {
			Response(c, http.StatusOK, errCode.SUCCESS, data)
			return
		}
	}
	Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
}

type UserProfileRequest struct {
	Uid      int    `form:"uid" json:"uid" biding:"required"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
}

// @Summary User Profile
// @Produce json
// @Param uid query int true "uid"
// @Router /api/v1/user/profile [get]
func GetUserProfile(c *gin.Context) {
	req := UserProfileRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}
	if err, data := models.UserProfile(req.Uid); err == nil {
		Response(c, http.StatusOK, errCode.SUCCESS, data)
	} else {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
	}
	Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
}

// @Summary User Profile
// @Produce json
// @Param uid query int true "uid"
// @Router /api/v1/user/profile [post]
func PostUserProfile(c *gin.Context) {
	req := UserProfileRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	if err := models.UpdateProfile(req); err == nil {
		Response(c, http.StatusOK, errCode.SUCCESS, nil)
	} else {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
	}
	Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
}
