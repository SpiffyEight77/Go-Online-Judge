package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"online-judge/controllers"
	_ "online-judge/docs"
	"online-judge/middlewares"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.GET("/login", controllers.GetUserLogin)
			user.POST("/login", controllers.PostUserLogin)
			user.GET("/register", controllers.GetUserRegister)
			user.POST("/register", controllers.PostUserRegister)
			user.GET("/profile", controllers.PostUserProfile).Use(middlewares.JWT())
			user.POST("/profile", controllers.PostUserProfile).Use(middlewares.JWT())
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
