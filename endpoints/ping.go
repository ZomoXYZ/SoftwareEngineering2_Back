package endpoints

import (
	"edu/letu/wan/database"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PingAuthorized(c *gin.Context) {
	if !database.IsAuthorized(c) {
		return
	}
	
	c.JSON(200, gin.H{
		"message": "pong",
	})
}