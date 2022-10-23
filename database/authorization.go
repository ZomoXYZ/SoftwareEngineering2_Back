package database

import (
	"edu/letu/wan/structs"

	"github.com/gin-gonic/gin"
)

func GetAuthorization(c *gin.Context) (string, *structs.PlayerInfo) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return "", nil
	}

	player := GetPlayerByToken(token)
	if player == nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return "", nil
	}

	//check if token is expired

	return token, player
}

func IsAuthorized(c *gin.Context) bool {
	token, _ := GetAuthorization(c)
	return token != ""
}
