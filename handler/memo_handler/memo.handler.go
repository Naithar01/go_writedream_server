package memoHandler

import (
	"github.com/Naithar01/go_write_dream/db"
	"github.com/Naithar01/go_write_dream/dto"
	"github.com/Naithar01/go_write_dream/models"
)

func GetAllMemoList() ([]models.MemoModel, error) {
	var memos []models.MemoModel

	// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
	rows, err := db.Database.Query("SELECT * FROM  memos")
	defer rows.Close()

	if err != nil {
		return []models.MemoModel{}, err
	}

	// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
	for rows.Next() {
		var memo models.MemoModel

		rows.Scan(&memo.Id, &memo.Issue_Id, &memo.Text, &memo.Created_At)

		memos = append(memos, memo)
	}

	return memos, nil
}

func CreateMemo(issue_id int, createMemoDTO dto.CreateMemoDTO) (int64, error) {
	// Body로 받은 Text와, Parmas로 받은 issueid의 값을 insert 해줌
	create_memo, err := db.Database.Exec("INSERT INTO  memos (issue_id, text) VALUES (?, ?)", issue_id, createMemoDTO.Text)

	if err != nil {
		return 0, err
	}

	created_memo_id, err := create_memo.LastInsertId()

	if err != nil {
		return 0, err
	}

	return created_memo_id, nil
}

func FindMemoById(id int) (models.MemoModel, error) {
	var memo models.MemoModel

	err := db.Database.QueryRow("SELECT * FROM  memos WHERE id = ?", id).Scan(&memo.Id, &memo.Issue_Id, &memo.Text, &memo.Created_At)

	if err != nil {
		return models.MemoModel{}, err
	}

	return memo, nil
}

func DeleteMemo(id int) (int, error) {
	_, err := db.Database.Exec("DELETE FROM  memos WHERE id = ?", id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
