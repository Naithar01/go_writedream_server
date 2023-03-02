package issueHandler

import (
	"errors"
	"fmt"

	"github.com/Naithar01/go_write_dream/db"
	"github.com/Naithar01/go_write_dream/dto"
	"github.com/Naithar01/go_write_dream/models"
)

func GetAllIssueList(issues_Query dto.IssueListQuery) ([]models.IssueListModel, int, error) {
	var issues []models.IssueListModel

	// Category Id Query가 있을 때
	// - 페이징 처리 Query가 없을 때
	// - 페이징 처리 Query가 있을 때
	if issues_Query.Category_Id >= 1 {
		rows, err := db.Database.Query("SELECT issue_id FROM  issue_category WHERE category_id = ?", issues_Query.Category_Id)
		defer rows.Close()

		if err != nil {
			return []models.IssueListModel{}, 0, err
		}

		var issue_id_list string

		for rows.Next() {
			var issue_id int
			rows.Scan(&issue_id)
			issue_id_list = fmt.Sprintf("%s'%d',", issue_id_list, issue_id)
		}

		if len(issue_id_list) == 0 {
			return []models.IssueListModel{}, 0, errors.New("검색된 Issue가 없습니다.")
		}

		issue_id_list = issue_id_list[:len(issue_id_list)-1]

		if issues_Query.Category_Id >= 1 && issues_Query.Page <= 0 && issues_Query.Page_Limit <= 0 { // Category Query가 있으면서, Page, Page_Limit Query가 없으면...
			sql := fmt.Sprintf("SELECT iss.id, iss.title, iss.content, iss.view_count, iss.create_at, iss.update_at, count(mms.id) AS memo_count from  issues AS iss LEFT OUTER JOIN  memos AS mms on iss.id = mms.issue_id WHERE iss.id in (%s) GROUP BY iss.id", issue_id_list)
			// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
			rows, err := db.Database.Query(sql)
			defer rows.Close()

			if err != nil {
				return []models.IssueListModel{}, 0, err
			}

			// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
			for rows.Next() {
				var issue models.IssueListModel

				rows.Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At, &issue.Memo_Count)

				issues = append(issues, issue)
			}

			return issues, len(issues), nil
		} else if issues_Query.Category_Id >= 1 && issues_Query.Page >= 1 && issues_Query.Page_Limit >= 1 { // Category Query가 있으면서, Page, Page_Limit Query가 둘 다 있으면...
			sql := fmt.Sprintf("SELECT iss.id, iss.title, iss.content, iss.view_count, iss.create_at, iss.update_at, count(mms.id) AS memo_count from  issues AS iss LEFT OUTER JOIN  memos AS mms on iss.id = mms.issue_id WHERE iss.id in (%s) GROUP BY iss.id limit %d, %d", issue_id_list, (issues_Query.Page-1)*issues_Query.Page_Limit, issues_Query.Page_Limit)

			// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
			// Limit을 사용하여 페이징 처리를 해줄건데, (Page -1) * Page_limit, Page_limit * Page
			rows, err := db.Database.Query(sql)
			defer rows.Close()

			if err != nil {
				return []models.IssueListModel{}, 0, err
			}

			// 몇 개의 Issue가 있는지 체크
			sql = fmt.Sprintf("SELECT count(iss.id) AS issue_count from  issues AS iss WHERE iss.id in (%s)", issue_id_list)
			var issue_count int
			err = db.Database.QueryRow(sql).Scan(&issue_count)

			if err != nil {
				return []models.IssueListModel{}, 0, err
			}

			// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
			for rows.Next() {
				var issue models.IssueListModel

				rows.Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At, &issue.Memo_Count)

				issues = append(issues, issue)
			}

			return issues, issue_count, nil
		}

	} else {
		// Category Id Query가 없을 때
		// - 페이징 처리 Query가 없을 때
		// - 페이징 처리 Query가 있을 때
		if issues_Query.Category_Id <= 0 && issues_Query.Page >= 1 && issues_Query.Page_Limit >= 1 { // Category Query가 없으면서, Page, Page_Limit Query가 있으면...
			sql := fmt.Sprintf("SELECT iss.id, iss.title, iss.content, iss.view_count, iss.create_at, iss.update_at, count(mms.id) AS memo_count from  issues AS iss LEFT OUTER JOIN  memos AS mms on iss.id = mms.issue_id GROUP BY iss.id limit %d, %d", (issues_Query.Page-1)*issues_Query.Page_Limit, issues_Query.Page_Limit)
			// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
			// Limit을 사용하여 페이징 처리를 해줄건데, (Page -1) * Page_limit, Page_limit * Page
			rows, err := db.Database.Query(sql)
			defer rows.Close()

			if err != nil {
				return []models.IssueListModel{}, 0, err
			}

			// 몇 개의 Issue가 있는지 체크
			sql = fmt.Sprintf("SELECT count(iss.id) AS issue_count from  issues AS iss")
			var issue_count int
			err = db.Database.QueryRow(sql).Scan(&issue_count)

			if err != nil {
				return []models.IssueListModel{}, 0, err
			}

			// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
			for rows.Next() {
				var issue models.IssueListModel

				rows.Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At, &issue.Memo_Count)

				issues = append(issues, issue)
			}

			return issues, issue_count, nil
		} else if issues_Query.Category_Id <= 0 && issues_Query.Page <= 0 && issues_Query.Page_Limit <= 0 { // Catgory Query가 없으면서, Page, Page_Limit Query가 없으면...
			sql := fmt.Sprintf("SELECT iss.id, iss.title, iss.content, iss.view_count, iss.create_at, iss.update_at, count(mms.id) AS memo_count from  issues AS iss LEFT OUTER JOIN  memos AS mms on iss.id = mms.issue_id GROUP BY iss.id")
			// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
			rows, err := db.Database.Query(sql)
			defer rows.Close()

			if err != nil {
				return []models.IssueListModel{}, 0, err
			}

			// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
			for rows.Next() {
				var issue models.IssueListModel

				rows.Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At, &issue.Memo_Count)

				issues = append(issues, issue)
			}

			return issues, len(issues), nil
		}
	}

	return []models.IssueListModel{}, 0, errors.New("올바르지 않은 요청입니다.")
}

func CreateIssue(issues_category dto.CreateIssueCategoryDTO, issue dto.CreateIssueDTO) (int64, error) {
	// Title과 Content를 Insert 해줌
	create_issue, err := db.Database.Exec("INSERT INTO  issues (title, content) VALUES (?, ?)", issue.Title, issue.Content)

	if err != nil {
		return 0, err
	}

	// 위에서 코드에서 테이블에 insert에 성공하였다면 새로운 행이 생겼다는 뜻이고
	// 마지막에 생긴 행이 될테니까 마지막 행의 Id 값을 가져오면 방금 생성했던 Issue의 Id를 가져올 수 있다.
	created_issue_id, err := create_issue.LastInsertId()

	if err != nil {
		return 0, err
	}

	_, err = db.Database.Exec("INSERT INTO  issue_category (issue_id, category_id) VALUES (?, ?)", created_issue_id, issues_category.Category_Id)

	if err != nil {
		return 0, err
	}

	return created_issue_id, nil
}

func FindIssueById(id int) (models.IssueFindModel, error) {
	var issue models.IssueFindModel

	// issue 테이블에서 id로 특정 행을 찾고 만약에 행이 존재하면 그 행의 값을 Scan하여 특정 값을 가져옴
	err := db.Database.QueryRow("SELECT * FROM  issues WHERE id = ?", id).Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At)

	if err != nil {
		return models.IssueFindModel{}, err
	}

	// 특정 issue를 검색할 때마다 issue의 view_count 열의 값을 1씩 올려줌
	_, err = db.Database.Exec("UPDATE  issues SET view_count = view_count + 1 WHERE id = ?", id)

	if err != nil {
		return models.IssueFindModel{}, err
	}

	// Memo 테이블에서 Issue의 id를 외래키로 저장하고 있는 행을 모두 가져옴
	rows, err := db.Database.Query("SELECT * FROM  memos WHERE issue_id = ?", id)
	defer rows.Close()

	if err != nil {
		return models.IssueFindModel{}, err
	}

	var memos []models.MemoModel

	// Rows에 SELECT한 행들이 모두 들어간다.
	// 각 행의 값을 Scan하여 Memos에 append (push) 해주고
	// 모든 행을 다 읽으면 반복문 탈출
	// memos를 issue에 Memos 컬럼에 넣어준다.
	for rows.Next() {
		var memo models.MemoModel

		rows.Scan(&memo.Id, &memo.Issue_Id, &memo.Text, &memo.Created_At)

		memos = append(memos, memo)
	}

	issue.Memos = memos

	return issue, nil
}

func UpdateIssue(id int, issue models.IssueModel) (int, error) {
	// Update Query를 사용하여 Issue 테이블의 Id에 맞는 raw를 Update 해줌
	_, err := db.Database.Exec("UPDATE  issues SET title = ?, content = ? WHERE id = ?", issue.Title, issue.Content, id)

	// 만약에 업데이트를 했는데 오류가 생겼다면 에러
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteIssue(id int) (int, error) {
	// Delete Query를 사용하여 Issue 테이블에 Id에 맞는 raw을 삭제해줌
	_, err := db.Database.Exec("DELETE FROM  issues WHERE id = ?", id)
	db.Database.Exec("DELETE FROM  issue_category WHERE issue_id = ? or category_id = ?", id, id)

	// Delete를 할 때 오류가 생겼다면...
	if err != nil {
		return 0, err
	}

	return id, nil
}
