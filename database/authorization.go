package database

import (
	"edu/letu/wan/structs"

	"github.com/gin-gonic/gin"
)

func GetAuthorization(c *gin.Context) (string, *structs.PlayerInfo) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(401, structs.ErrorJson{Error: "Unauthorized"})
		return "", nil
	}

	player := GetPlayerByToken(token)
	if player == nil {
		c.AbortWithStatusJSON(401, structs.ErrorJson{Error: "Unauthorized"})
		return "", nil
	}

	return token, player
}

func IsAuthorized(c *gin.Context) bool {
	token, _ := GetAuthorization(c)
	return token != ""
}
