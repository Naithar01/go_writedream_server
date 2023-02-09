package router

import (
	issueHandler "github.com/Naithar01/go_write_dream/issue_handler"
	"github.com/Naithar01/go_write_dream/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	app := gin.Default()
	app.Use(cors.Default())

	api := app.Group("/api")
	api.Use(middleware.SetHader)
	{
		issue := api.Group("/issues")
		{
			issue.GET("/", issueHandler.GetAllIssueList)
			issue.GET("/:id", issueHandler.FindIssueById)
			issue.POST("/", issueHandler.CreateIssue)
			issue.PATCH("/:id", issueHandler.UpdateIssue)
			issue.PUT("/:id", issueHandler.UpdateIssue)
			issue.DELETE("/:id", issueHandler.DeleteIssue)
		}

	}

	return app
}
