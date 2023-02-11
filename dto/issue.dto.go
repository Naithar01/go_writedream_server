package dto

type CreateIssueDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateIssueDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type IssueListQuery struct {
	Category_Id int `json:"category_id" query:"category_id" form:"category_id"`
	Page        int `json:"page" query:"page" form:"page"`
	Page_Limit  int `json:"Page_Limit" query:"page_limit" form:"page_limit"`
}

type CreateIssueCategoryDTO struct {
	Category_Id int `json:"category_id" query:"category_id" form:"category_id"`
}
