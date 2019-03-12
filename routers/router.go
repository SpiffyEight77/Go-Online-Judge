package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"online-judge/controllers"
	_ "online-judge/docs"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.GET("/login", controllers.GetUserLogin)
			user.POST("/login", controllers.PostUserLogin)
			user.GET("/register", controllers.GetUserRegister)
			user.POST("/register", controllers.PostUserRegister)

			//profile := user.Group("/profile").Use(middlewares.JWT())
			profile := user.Group("/profile")
			{
				profile.GET("/detail", controllers.GetUserProfile)
				profile.POST("/edit", controllers.PostUserProfile)
			}
		}
		problem := v1.Group("/problem")
		{
			problem.GET("/list", controllers.GetProblemList)
			problem.GET("/detail", controllers.GetProblemDetail)
			problem.POST("/submit", controllers.PostSubmitProblem)
			problem.POST("/create", controllers.PostCreateProblem)
			problem.POST("/update", controllers.PostUpdateProblem)
			problem.POST("/delete", controllers.PostDeleteProblem)
		}
		news := v1.Group("/news")
		{
			news.GET("/list", controllers.GetNewsList)
			news.GET("/detail", controllers.GetNewsDetail)
		}
		submission := v1.Group("/submission")
		{
			submission.POST("/submit", controllers.PostSubmission)
			submission.GET("/list", controllers.GetSubmission)
			submission.GET("/solved", controllers.GetSolvedProblems)
		}
		contest := v1.Group("/contest")
		{
			contest.GET("/list", controllers.GetContestList)
			contest.GET("/detail", controllers.GetContestDetail)
			contest.GET("/submission", controllers.GetContestSubmission)
			contest.POST("/update",controllers)
			//contest.POST("/submit", controllers.PostContestProblemSubmit)
			problem := contest.Group("/problem")
			{
				problem.GET("/detail", controllers.GetContestProblemDetail)
				problem.POST("/submit", controllers.PostContestProblemSubmit)
			}
		}
		solution := v1.Group("/solution")
		{
			solution.GET("/list", controllers.GetSolutionList)
			solution.GET("/detail", controllers.GetSolutionDetail)
		}

		administration := v1.Group("/admin")
		{
			user := administration.Group("/user")
			{
				user.GET("/list", controllers.GetUserList)
				user.GET("/login", controllers.GetUserLogin)
				user.POST("/login", controllers.PostUserLogin)
				//user.POST("/delete", controllers.PostDeleteUser)
			}
			problem := administration.Group("/problem")
			{
				problem.GET("/list", controllers.GetProblemList)
				problem.GET("/detail", controllers.GetProblemDetail)
				//problem.POST("/delete", controllers.PostDeleteProblem)
				//problem.POST("/new", controllers.PostCreateProblem)
				//problem.POST("/edit", controllers.PostUpdateProblem)
			}
			news := administration.Group("/news")
			{
				news.GET("/list", controllers.GetNewsList)
				news.GET("/detail", controllers.GetNewsDetail)
				news.POST("/edit", controllers.PostNewsEdit)
				news.POST("/create", controllers.PostNewsCreate)
				news.POST("/delete", controllers.PostNewsDelete)
			}
			contest := administration.Group("/contest")
			{
				contest.GET("/list", controllers.GetContestList)
				contest.GET("/detail", controllers.GetContestDetail)
				contest.POST("/create", controllers.PostCreateContest)
				//contest.POST("/edit", controllers.PostUpdateContest)
				contest.POST("/delete", controllers.PostDeleteContest)
			}
			solution := administration.Group("/solution")
			{
				solution.GET("/list", controllers.GetSolutionList)
				solution.GET("/detail", controllers.GetSolutionDetail)
			}
		}
	}
	return router
}
