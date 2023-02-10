package dto

type CreateIssueDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateIssueDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
