package memoHandler

import (
	"net/http"

	"github.com/Naithar01/go_write_dream/db"
	"github.com/Naithar01/go_write_dream/dto"
	errorHandler "github.com/Naithar01/go_write_dream/handler/error_handler"
	"github.com/Naithar01/go_write_dream/models"
	"github.com/gin-gonic/gin"
)

func GetAllMemoList(c *gin.Context) {
	var memos []models.MemoModel

	// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
	rows, err := db.Database.Query("SELECT * FROM writedream.memos")
	defer rows.Close()

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
	for rows.Next() {
		var memo models.MemoModel

		rows.Scan(&memo.Id, &memo.Issue_Id, &memo.Text, &memo.Created_At)

		memos = append(memos, memo)
	}

	c.JSON(http.StatusOK, gin.H{
		"memos": memos,
	})
}

func CreateMemo(c *gin.Context) {
	issue_id := c.Param("issueid")

	// Parmas로 받은 id가 없다면 에러 반환
	if len(issue_id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Params로 받아야하는 id가 없습니다.",
		})
		return
	}

	var memo dto.CreateMemoDTO

	if err := c.BindJSON(&memo); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// 만약 Body에 들어온 Text가 빈 문자열이라면 에러 반환
	if len(memo.Text) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Text의 문자 길이가 너무 짧습니다.",
		})
		return
	}

	// Body로 받은 Text와, Parmas로 받은 issueid의 값을 insert 해줌
	create_memo, err := db.Database.Exec("INSERT INTO writedream.memos (issue_id, text) VALUES (?, ?)", issue_id, memo.Text)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	created_memo_id, err := create_memo.LastInsertId()

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": created_memo_id,
	})
}

func FindMemoById(c *gin.Context) {
	id := c.Param("id")

	// Parmas로 받은 id가 없다면 에러 반환
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Params로 받아야하는 id가 없습니다.",
		})
		return
	}

	var memo models.MemoModel

	err := db.Database.QueryRow("SELECT * FROM writedream.memos WHERE id = ?", id).Scan(&memo.Id, &memo.Issue_Id, &memo.Text, &memo.Created_At)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"memo": memo,
	})
}

func DeleteMemo(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "id가 없으면 검색할 수 없습니다.",
		})
		return
	}

	_, err := db.Database.Exec("DELETE FROM writedream.memos WHERE id = ?", id)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
