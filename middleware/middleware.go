package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SetHader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0")
	c.Header("Last-Modified", time.Now().String())

	c.Next()
}
