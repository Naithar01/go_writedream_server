package router

import (
	"github.com/Naithar01/go_write_dream/controllers"
	"github.com/Naithar01/go_write_dream/middleware"
	"github.com/gin-gonic/gin"
)

var (
	CategoryController controllers.CategoryController = *controllers.NewCategoryController()
	IssueController    controllers.IssueController    = *controllers.NewIssueController()
	MemoController     controllers.MemoController     = *controllers.NewMemoController()
)

func InitRouter() *gin.Engine {
	app := gin.Default()
	app.Use(middleware.CORSMiddleware())
	api := app.Group("/api")
	{
		categories := api.Group("/categories")
		{
			categories.GET("/", CategoryController.GetAllCategoryList)
			categories.POST("/", CategoryController.CreateCategory)
			categories.DELETE("/:id", CategoryController.DeleteCategory)
		}
		issue := api.Group("/issues")
		{
			issue.GET("/", IssueController.GetAllIssueList)
			issue.GET("/:id", IssueController.FindIssueById)
			issue.POST("/", IssueController.CreateIssue)
			issue.PATCH("/:id", IssueController.UpdateIssue)
			issue.PUT("/:id", IssueController.UpdateIssue)
			issue.DELETE("/:id", IssueController.DeleteIssue)
		}
		memo := issue.Group("/memos")
		{
			memo.GET("/", MemoController.GetAllMemoList)
			memo.GET("/:id", MemoController.FindMemoById)
			memo.POST("/:issueid", MemoController.CreateMemo)
			memo.DELETE("/:id", MemoController.DeleteMemo)
		}
	}

	return app
}
