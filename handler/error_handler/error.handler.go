package errorHandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err error) {
	log.Println(err.Error())
	c.JSON(http.StatusBadRequest, gin.H{
		"Error": err.Error(),
	})
}
