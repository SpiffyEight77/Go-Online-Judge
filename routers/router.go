package routers

import (
	"github.com/gin-gonic/gin"
	"online-judge/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", controllers.Home)
	router.GET("/faq", controllers.Faq)
	router.GET("/problem", controllers.Problem)
	router.GET("/login", controllers.Login)
	router.GET("/register", controllers.Register)
	router.POST("/register", controllers.PostRegister)


	return router
}
