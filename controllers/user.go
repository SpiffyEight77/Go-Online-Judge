package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"online-judge/common/errCode"
	"online-judge/middlewares/jwt"
	"online-judge/models"
	"strconv"
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
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if ok {
		if isExist, data := user.CheckAuth(); isExist == true {
			if token, err := jwt.GenerateToken(req.Username, req.Password); err != nil {
				code = errCode.UNAUTHORIZED
			} else {
				data.Token = token
				code = errCode.SUCCESS
				user.Token = token
				user.LastLogin = time.Now()
				user.UpdateUserLogin()
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
// @Router /api/v1/user/register [get]
func GetUserRegister(c *gin.Context) {
	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

// @Summary User Register
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {string} json "{"data":{"token":(string),"username":(string)},"msg":"success"}"
// @Router /api/v1/user/register [post]
func PostUserRegister(c *gin.Context) {
	req := UserRegisterRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	user := models.User{
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		CreatedAt: time.Now(),
		LastLogin: time.Now(),
	}

	ok, data := user.CheckAuth()
	if ok || data != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	token, err := jwt.GenerateToken(req.Username, req.Password)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	user.Token = token
	if err := user.Register(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	ok, data = user.CheckAuth()
	if !ok {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

type UserProfileRequest struct {
	Uid      int    `form:"uid" json:"uid" biding:"required"`
	Username string `form:"username" json:"username" biding:"-"`
	Password string `form:"password" json:"password" biding:"-"`
	Email    string `form:"email" json:"email" biding:"-"`
}

// @Summary User Profile
// @Produce json
// @Param uid query int true "uid"
// @Success 200 {string} json "{"data":{"token":(string),"username":(string)},"msg":"success"}"
// @Router /api/v1/user/profile/detail [get]
func GetUserProfile(c *gin.Context) {
	id := c.GetHeader("uid")
	if id == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	uid, err := strconv.Atoi(id)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	user := models.User{
		ID: uid,
	}
	data, err := user.UserProfile()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

// @Summary User Profile
// @Produce json
// @Param uid query int true "uid"
// @Router /api/v1/user/profile/detail [post]
func PostUserProfile(c *gin.Context) {
	req := UserProfileRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	user := models.User{
		ID:       req.Uid,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	if err := user.UpdateProfile(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

type UserDeleteRequest struct {
	IDList []int `form:"id_list" json:"id_list" biding:"required"`
}

// @Summary Delete User
// @Produce json
// @Param id_list query int true "id_list"
// @Router /api/v1/user/delete [post]
func PostDeleteUser(c *gin.Context) {
	req := UserDeleteRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	user := models.User{
		IDList: req.IDList,
	}
	if err := user.DeleteUser(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

// @Summary User List
// @Produce json
// @Router /api/v1/user/list [get]
func GetUserList(c *gin.Context) {
	var userList models.User
	data, err := userList.UserList()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}
