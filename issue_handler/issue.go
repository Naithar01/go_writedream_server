package issueHandler

import (
	"log"
	"net/http"

	"github.com/Naithar01/go_write_dream/db"
	"github.com/Naithar01/go_write_dream/models"
	"github.com/gin-gonic/gin"
)

// Get Select All issue

func GetAllIssueList(c *gin.Context) {
	var issues []models.IssueModel

	rows, err := db.Database.Query("SELECT * FROM writedream.issues")

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	}

	for rows.Next() {
		var issue models.IssueModel

		rows.Scan(&issue.Id, &issue.Title, &issue.Content, &issue.ViewCount, &issue.Created_At, &issue.Updated_At)

		issues = append(issues, issue)
	}

	c.JSON(http.StatusOK, gin.H{
		"Issues": issues,
	})
}
