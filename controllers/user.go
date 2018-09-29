package controllers

import "github.com/gin-gonic/gin"

type UserLoginRequest struct {
	Username string `form:"username" json:"username" biding:"required"`
	Password string `form:"password" json:"password" biding:"required"`
}

// @Summary Admin Login
// @Produce json
// @Router /api/v1/admin/user/login [get]
func GetUserLogin(c *gin.Context) {

}

// @Summary Admin Login
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Router /api/v1/admin/user/login [post]
func PostUserLogin(c *gin.Context) {

}

// @Summary User Register
// @Produce json
// @Router /api/v1/user/login [get]
func GetUserRegister(c *gin.Context) {

}

// @Summary User Login
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Router /api/v1/user/login [post]
func PostUserRegister(c *gin.Context) {

}

// @Summary User Profile
// @Produce json
// @Router /api/v1/user/profile [get]
func GetUserProfile(c *gin.Context) {

}

// @Summary User Login
// @Produce json
// @Param uid query int true "uid"
// @Router /api/v1/user/profile [post]
func PostUserProfile(c *gin.Context) {

}
