package models

import "time"

type IssueModel struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	ViewCount  int       `json:"view_count"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type IssueListModel struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	ViewCount  int       `json:"view_count"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
	Memo_Count int       `json:"memo_count"`
}
