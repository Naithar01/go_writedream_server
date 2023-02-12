package issueHandler

import (
	"fmt"
	"net/http"

	"github.com/Naithar01/go_write_dream/db"
	"github.com/Naithar01/go_write_dream/dto"
	errorHandler "github.com/Naithar01/go_write_dream/handler/error_handler"
	"github.com/Naithar01/go_write_dream/models"
	"github.com/gin-gonic/gin"
)

func GetAllIssueList(c *gin.Context) {
	var issues []models.IssueListModel

	var issues_Query dto.IssueListQuery

	// Query Parameter로 들어오는 Page, Page_limit의 값을 각 변수에 담아줌
	if err := c.BindQuery(&issues_Query); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// Category Id Query가 있을 때
	// - 페이징 처리 Query가 없을 때
	// - 페이징 처리 Query가 있을 때
	if issues_Query.Category_Id >= 1 {
		rows, err := db.Database.Query("SELECT issue_id FROM writedream.issue_category WHERE category_id = ?", issues_Query.Category_Id)
		defer rows.Close()

		if err != nil {
			errorHandler.ErrorHandler(c, err)
			return
		}

		var issue_id_list string

		for rows.Next() {
			var issue_id int
			rows.Scan(&issue_id)
			issue_id_list = fmt.Sprintf("%s'%d',", issue_id_list, issue_id)
		}

		issue_id_list = issue_id_list[:len(issue_id_list)-1]

		if issues_Query.Category_Id >= 1 && issues_Query.Page <= 0 && issues_Query.Page_Limit <= 0 { // Category Query가 있으면서, Page, Page_Limit Query가 없으면...
			sql := fmt.Sprintf("SELECT iss.id, iss.title, iss.content, iss.view_count, iss.create_at, iss.update_at, count(mms.id) AS memo_count from writedream.issues AS iss LEFT OUTER JOIN writedream.memos AS mms on iss.id = mms.issue_id WHERE iss.id in (%s) GROUP BY iss.id", issue_id_list)
			// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
			rows, err := db.Database.Query(sql)
			defer rows.Close()

			if err != nil {
				errorHandler.ErrorHandler(c, err)
				return
			}

			// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
			for rows.Next() {
				var issue models.IssueListModel

				rows.Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At, &issue.Memo_Count)

				issues = append(issues, issue)
			}

			c.JSON(http.StatusOK, gin.H{
				"issues": issues,
			})
			return
		} else if issues_Query.Category_Id >= 1 && issues_Query.Page >= 1 && issues_Query.Page_Limit >= 1 { // Category Query가 있으면서, Page, Page_Limit Query가 둘 다 있으면...
			sql := fmt.Sprintf("SELECT iss.id, iss.title, iss.content, iss.view_count, iss.create_at, iss.update_at, count(mms.id) AS memo_count from writedream.issues AS iss LEFT OUTER JOIN writedream.memos AS mms on iss.id = mms.issue_id WHERE iss.id in (%s) GROUP BY iss.id limit %d, %d", issue_id_list, (issues_Query.Page-1)*issues_Query.Page_Limit, issues_Query.Page_Limit)

			// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
			// Limit을 사용하여 페이징 처리를 해줄건데, (Page -1) * Page_limit, Page_limit * Page
			rows, err := db.Database.Query(sql)
			defer rows.Close()

			if err != nil {
				errorHandler.ErrorHandler(c, err)
				return
			}

			// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
			for rows.Next() {
				var issue models.IssueListModel

				rows.Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At, &issue.Memo_Count)

				issues = append(issues, issue)
			}

			c.JSON(http.StatusOK, gin.H{
				"issues": issues,
			})
			return
		}

	} else {
		// Category Id Query가 없을 때
		// - 페이징 처리 Query가 없을 때
		// - 페이징 처리 Query가 있을 때
		if issues_Query.Category_Id <= 0 && issues_Query.Page >= 1 && issues_Query.Page_Limit >= 1 { // Category Query가 없으면서, Page, Page_Limit Query가 있으면...
			sql := fmt.Sprintf("SELECT iss.id, iss.title, iss.content, iss.view_count, iss.create_at, iss.update_at, count(mms.id) AS memo_count from writedream.issues AS iss LEFT OUTER JOIN writedream.memos AS mms on iss.id = mms.issue_id GROUP BY iss.id limit %d, %d", (issues_Query.Page-1)*issues_Query.Page_Limit, issues_Query.Page_Limit)
			// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
			// Limit을 사용하여 페이징 처리를 해줄건데, (Page -1) * Page_limit, Page_limit * Page
			rows, err := db.Database.Query(sql)
			defer rows.Close()

			if err != nil {
				errorHandler.ErrorHandler(c, err)
				return
			}

			// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
			for rows.Next() {
				var issue models.IssueListModel

				rows.Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At, &issue.Memo_Count)

				issues = append(issues, issue)
			}

			c.JSON(http.StatusOK, gin.H{
				"issues": issues,
			})
			return
		} else if issues_Query.Category_Id <= 0 && issues_Query.Page <= 0 && issues_Query.Page_Limit <= 0 { // Catgory Query가 없으면서, Page, Page_Limit Query가 없으면...
			sql := fmt.Sprintf("SELECT iss.id, iss.title, iss.content, iss.view_count, iss.create_at, iss.update_at, count(mms.id) AS memo_count from writedream.issues AS iss LEFT OUTER JOIN writedream.memos AS mms on iss.id = mms.issue_id GROUP BY iss.id")
			// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
			rows, err := db.Database.Query(sql)
			defer rows.Close()

			if err != nil {
				errorHandler.ErrorHandler(c, err)
				return
			}

			// rows 변수를 한 행씩 읽어내려가는데 마지막 행을 읽고 다음 행은 없는 행이 되니까 false를 return, 반복문이 끝남
			for rows.Next() {
				var issue models.IssueListModel

				rows.Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At, &issue.Memo_Count)

				issues = append(issues, issue)
			}

			c.JSON(http.StatusOK, gin.H{
				"issues": issues,
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"Error": "Error",
	})
	return
}

func CreateIssue(c *gin.Context) {
	var issues_category dto.CreateIssueCategoryDTO

	// Query Parameter로 들어오는 category를 issues_category 변수에 담아줌
	if err := c.BindQuery(&issues_category); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// Query로 들어온 category_id가 비어있으면 에러 반환
	if issues_category.Category_Id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "존재하지 않은 Category_id 입니다.",
		})
		return
	}

	var issue dto.CreateIssueDTO

	if err := c.BindJSON(&issue); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// 만약 Body에 들어온 Title 혹은 Content가 빈 문자열이라면 에러 반환
	if len(issue.Title) == 0 || len(issue.Content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Title 혹은 Content의 문자 길이가 너무 짧습니다.",
		})
		return
	}

	// Title과 Content를 Insert 해줌
	create_issue, err := db.Database.Exec("INSERT INTO writedream.issues (title, content) VALUES (?, ?)", issue.Title, issue.Content)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// 위에서 코드에서 테이블에 insert에 성공하였다면 새로운 행이 생겼다는 뜻이고
	// 마지막에 생긴 행이 될테니까 마지막 행의 Id 값을 가져오면 방금 생성했던 Issue의 Id를 가져올 수 있다.
	created_issue_id, err := create_issue.LastInsertId()

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	_, err = db.Database.Exec("INSERT INTO writedream.issue_category (issue_id, category_id) VALUES (?, ?)", created_issue_id, issues_category.Category_Id)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": created_issue_id,
	})
}

func FindIssueById(c *gin.Context) {
	id := c.Param("id")

	var issue models.IssueFindModel

	// issue 테이블에서 id로 특정 행을 찾고 만약에 행이 존재하면 그 행의 값을 Scan하여 특정 값을 가져옴
	err := db.Database.QueryRow("SELECT * FROM writedream.issues WHERE id = ?", id).Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// 특정 issue를 검색할 때마다 issue의 view_count 열의 값을 1씩 올려줌
	_, err = db.Database.Exec("UPDATE writedream.issues SET view_count = view_count + 1 WHERE id = ?", id)

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// Memo 테이블에서 Issue의 id를 외래키로 저장하고 있는 행을 모두 가져옴
	rows, err := db.Database.Query("SELECT * FROM writedream.memos WHERE issue_id = ?", id)
	defer rows.Close()

	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
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

	c.JSON(http.StatusOK, gin.H{
		"issue": issue,
	})
}

func UpdateIssue(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "id가 없으면 검색할 수 없습니다.",
		})
		return
	}

	var issue models.IssueModel

	// Body로 들어오는 JSON 값을 UpdateIssueDTO 객체 값에 맞게 매핑
	if err := c.BindJSON(&issue); err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	// 만약 Body에 들어온 Title 혹은 Content가 빈 문자열이라면 에러 반환
	if len(issue.Title) == 0 || len(issue.Content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Title 혹은 Content의 문자 길이가 너무 짧습니다.",
		})
		return
	}

	// Update Query를 사용하여 Issue 테이블의 Id에 맞는 raw를 Update 해줌
	_, err := db.Database.Exec("UPDATE writedream.issues SET title = ?, content = ? WHERE id = ?", issue.Title, issue.Content, id)

	// 만약에 업데이트를 했는데 오류가 생겼다면 에러
	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func DeleteIssue(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "id가 없으면 검색할 수 없습니다.",
		})
		return
	}

	// Delete Query를 사용하여 Issue 테이블에 Id에 맞는 raw을 삭제해줌
	_, err := db.Database.Exec("DELETE FROM writedream.issues WHERE id = ?", id)
	db.Database.Exec("DELETE FROM writedream.issue_category WHERE issue_id = ? or category_id = ?", id, id)

	// Delete를 할 때 오류가 생겼다면...
	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
