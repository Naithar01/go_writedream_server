package categoryhandler

import (
	"github.com/Naithar01/go_write_dream/db"
	"github.com/Naithar01/go_write_dream/dto"
	"github.com/Naithar01/go_write_dream/models"
)

func GetAllCategoryList() ([]models.CategoryModel, error) {
	var categories []models.CategoryModel

	// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
	rows, err := db.Database.Query("SELECT * FROM  categories;")
	defer rows.Close()

	if err != nil {
		return []models.CategoryModel{}, err
	}

	// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
	for rows.Next() {
		var category models.CategoryModel

		rows.Scan(&category.Id, &category.Title)

		categories = append(categories, category)
	}

	return categories, nil
}

func CreateCategory(createCategoryDTO dto.CreateCategoryDTO) (int64, error) {
	// Body로 받은 Title을 insert 해줌
	create_category, err := db.Database.Exec("INSERT INTO  categories (title) VALUES (?)", createCategoryDTO.Title)

	if err != nil {
		return 0, err
	}

	created_category_id, err := create_category.LastInsertId()

	if err != nil {
		return 0, err
	}

	return created_category_id, nil
}

func DeleteCategory(id int) (int, error) {
	_, err := db.Database.Exec("DELETE FROM  categories WHERE id = ?", id)
	db.Database.Exec("DELETE FROM  issue_category WHERE id = ?", id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
