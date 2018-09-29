package routers

import (
	"github.com/gin-gonic/gin"
	"online-judge/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()



	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.GET("/login", controllers.GetUserLogin)
			user.POST("/login", controllers.PostUserLogin)
			user.GET("/register", controllers.GetUserRegister)
			user.POST("/register", controllers.PostUserRegister)
			user.GET("/profile")
			user.POST("/profile")
		}
		administration := v1.Group("/admin")
		{
			user := administration.Group("/user")
			{
				user.GET("/login", controllers.GetUserLogin)
				user.POST("/login", controllers.PostUserLogin)
			}
		}
	}
	return router
}
