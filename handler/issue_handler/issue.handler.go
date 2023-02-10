package issueHandler

import (
	"net/http"

	"github.com/Naithar01/go_write_dream/db"
	"github.com/Naithar01/go_write_dream/dto"
	errorHandler "github.com/Naithar01/go_write_dream/handler/error_handler"
	"github.com/Naithar01/go_write_dream/models"
	"github.com/gin-gonic/gin"
)

func GetAllIssueList(c *gin.Context) {
	var issues []models.IssueListModel

	// DB에서 SELECT 해온 모든 데이터들이 rows 변수에 담김
	rows, err := db.Database.Query("select iss.id, iss.title, iss.content, iss.view_count, iss.create_at, iss.update_at, count(mms.id) AS memo_count from writedream.issues AS iss inner join writedream.memos AS mms on iss.id = mms.issue_id GROUP BY iss.id")

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
}

func CreateIssue(c *gin.Context) {
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

	c.JSON(http.StatusCreated, gin.H{
		"id": created_issue_id,
	})
}

func FindIssueById(c *gin.Context) {
	id := c.Param("id")

	var issue models.IssueModel

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

	// Delete를 할 때 오류가 생겼다면...
	if err != nil {
		errorHandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
