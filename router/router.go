package router

import (
	issueHandler "github.com/Naithar01/go_write_dream/issue_handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	app := gin.Default()
	app.Use(cors.Default())

	issue := app.Group("/api/issues")
	{
		issue.GET("/", issueHandler.GetAllIssueList)
		issue.GET("/:id", issueHandler.FindIssueById)
		issue.POST("/", issueHandler.CreateIssue)
	}

	return app
}
