package models

import "time"

type MemoModel struct {
	Id         int       `json:"id"`
	Issue_Id   int       `json:"issue_id"`
	Text       string    `json:"text"`
	Created_At time.Time `json:"created_at"`
}
