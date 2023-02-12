package controllers

import (
	"net/http"
	"strconv"

	"github.com/Naithar01/go_write_dream/dto"
	categoryhandler "github.com/Naithar01/go_write_dream/handler/category_handler"
	errorHandler "github.com/Naithar01/go_write_dream/handler/error_handler"
	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (cc *CategoryController) GetAllCategoryList(c *gin.Context) {
	categories, err := categoryhandler.GetAllCategoryList()

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var createCategoryDTO dto.CreateCategoryDTO

	if err := c.BindJSON(&createCategoryDTO); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	if len(createCategoryDTO.Title) == 0 {
		c.JSON(http.StatusCreated, gin.H{
			"Error": "Title은 빈 문자열이 될 수 없습니다.",
		})
		return
	}

	created_category_id, err := categoryhandler.CreateCategory(createCategoryDTO)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": created_category_id,
	})
}

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "id가 없으면 검색할 수 없습니다.",
		})
		return
	}

	deleted_id, err := categoryhandler.DeleteCategory(id)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": deleted_id,
	})

}
