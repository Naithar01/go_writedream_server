package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Naithar01/go_write_dream/dto"
	errorHandler "github.com/Naithar01/go_write_dream/handler/error_handler"
	issueHandler "github.com/Naithar01/go_write_dream/handler/issue_handler"
	"github.com/Naithar01/go_write_dream/models"
	"github.com/gin-gonic/gin"
)

type IssueController struct{}

func NewIssueController() *IssueController {
	return &IssueController{}
}

func (ic *IssueController) GetAllIssueList(c *gin.Context) {
	var issues_Query dto.IssueListQuery

	// Query Parameter 값을 DTO에 맞게 넣어줌
	if err := c.BindQuery(&issues_Query); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	issues, issues_count, err := issueHandler.GetAllIssueList(issues_Query)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"issues":       issues,
		"issues_count": issues_count,
	})
}

func (ic *IssueController) CreateIssue(c *gin.Context) {
	var issues_category dto.CreateIssueCategoryDTO

	// Query Parameter로 들어오는 category를 issues_category 변수에 담아줌
	if err := c.BindQuery(&issues_category); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// Query로 들어온 category_id가 비어있으면 에러 반환
	if issues_category.Category_Id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "존재하지 않은 Category_id 입니다.",
		})
		return
	}

	var issue dto.CreateIssueDTO

	if err := c.BindJSON(&issue); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// 만약 Body에 들어온 Title 혹은 Content가 빈 문자열이라면 에러 반환
	if len(issue.Title) == 0 || len(issue.Content) == 0 {
		errorHandler.ErrorHandler(c, errors.New("Title 혹은 Content의 문자 길이가 너무 짧습니다."))
		return
	}

	created_id, err := issueHandler.CreateIssue(issues_category, issue)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": created_id,
	})
}

func (ic *IssueController) FindIssueById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	issue, err := issueHandler.FindIssueById(id)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"issue": issue,
	})

}

func (ic *IssueController) UpdateIssue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "id가 없으면 검색할 수 없습니다.",
		})
		return
	}

	var issue models.IssueModel

	// Body로 들어오는 JSON 값을 UpdateIssueDTO 객체 값에 맞게 매핑
	if err := c.BindJSON(&issue); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// 만약 Body에 들어온 Title 혹은 Content가 빈 문자열이라면 에러 반환
	if len(issue.Title) == 0 || len(issue.Content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Title 혹은 Content의 문자 길이가 너무 짧습니다.",
		})
		return
	}

	updated_id, err := issueHandler.UpdateIssue(id, issue)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": updated_id,
	})
}

func (ic *IssueController) DeleteIssue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "id가 없으면 검색할 수 없습니다.",
		})
		return
	}

	deleted_id, err := issueHandler.DeleteIssue(id)

	// Delete를 할 때 오류가 생겼다면...
	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": deleted_id,
	})
}
