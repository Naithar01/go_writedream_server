package categoryhandler

import (
	"net/http"

	"github.com/Naithar01/go_write_dream/db"
	"github.com/Naithar01/go_write_dream/dto"
	errorHandler "github.com/Naithar01/go_write_dream/handler/error_handler"
	"github.com/Naithar01/go_write_dream/models"
	"github.com/gin-gonic/gin"
)

func GetAllCategoryList(c *gin.Context) {
	var categories []models.CategoryModel

	// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
	rows, err := db.Database.Query("SELECT * FROM writedream.categories;")

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
	for rows.Next() {
		var category models.CategoryModel

		rows.Scan(&category.Id, &category.Title)

		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func CreateCategory(c *gin.Context) {
	var category dto.CreateCategoryDTO

	if err := c.BindJSON(&category); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// Body로 받은 Title을 insert 해줌
	create_category, err := db.Database.Exec("INSERT INTO writedream.categories (title) VALUES (?)", category.Title)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	created_category_id, err := create_category.LastInsertId()

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": created_category_id,
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "id가 없으면 검색할 수 없습니다.",
		})
		return
	}

	_, err := db.Database.Exec("DELETE FROM writedream.categories WHERE id = ?", id)
	db.Database.Exec("DELETE FROM writedream.issue_category WHERE id = ?", id)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
