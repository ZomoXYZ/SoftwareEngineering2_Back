package database

import "github.com/gin-gonic/gin"

func IsAuthorized(c *gin.Context) bool {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return false
	}

	player := GetPlayerByToken(token)
	if player == nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return false
	}

	//check if token is expired

	return true
}
