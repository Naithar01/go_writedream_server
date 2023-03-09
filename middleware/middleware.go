package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetHader(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PATCH, PUT, OPTIONS")

	c.Next()
}
