package controllers

import (
	"net/http"
	"strconv"

	"github.com/Naithar01/go_write_dream/dto"
	errorHandler "github.com/Naithar01/go_write_dream/handler/error_handler"
	memoHandler "github.com/Naithar01/go_write_dream/handler/memo_handler"
	"github.com/gin-gonic/gin"
)

type MemoController struct{}

func NewMemoController() *MemoController {
	return &MemoController{}
}

func (mc *MemoController) GetAllMemoList(c *gin.Context) {
	memos, err := memoHandler.GetAllMemoList()

	if err != nil {
		errorHandler.ErrorHandler(c, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"memos": memos,
	})
}

func (mc *MemoController) CreateMemo(c *gin.Context) {
	issue_id, _ := strconv.Atoi(c.Param("issueid"))

	if issue_id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Issue id가 잘못되어 Memo를 생성할 수 없습니다.",
		})
		return
	}

	var createMemoDTO dto.CreateMemoDTO

	if err := c.BindJSON(&createMemoDTO); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// 만약 Body에 들어온 Text가 빈 문자열이라면 에러 반환
	if len(createMemoDTO.Text) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Text의 문자 길이가 너무 짧습니다.",
		})
		return
	}

	created_memo_id, err := memoHandler.CreateMemo(issue_id, createMemoDTO)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": created_memo_id,
	})
}

func (mc *MemoController) FindMemoById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Parmas로 받은 id가 없다면 에러 반환
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Params로 받아야하는 id가 없습니다.",
		})
		return
	}

	memo, err := memoHandler.FindMemoById(id)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"memo": memo,
	})
}

func (mc *MemoController) DeleteMemo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Parmas로 받은 id가 없다면 에러 반환
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Params로 받아야하는 id가 없습니다.",
		})
		return
	}

	deleted_id, err := memoHandler.DeleteMemo(id)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": deleted_id,
	})
}
