package dto

type CreateIssueDTO struct {
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

type UpdateIssueDTO struct {
	Title   string `json:"Title"`
	Content string `json:"Content"`
}
