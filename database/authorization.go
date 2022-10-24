package database

import (
	"edu/letu/wan/structs"

	"github.com/gin-gonic/gin"
)

func GetAuthHeaders(c *gin.Context) (string, string) {
	token := c.Request.Header.Get("Authorization")
	uuid := c.Request.Header.Get("UUID")

	return token, uuid
}

func GetAuthorization(c *gin.Context) (string, string, *structs.PlayerInfo) {
	token, uuid := GetAuthHeaders(c)
	if token == "" || uuid == "" {
		c.AbortWithStatusJSON(401, structs.ErrorJson{Error: "Unauthorized"})
		return "", "", nil
	}

	exists := AuthorizationExists(token, uuid)
	if !exists {
		c.AbortWithStatusJSON(401, structs.ErrorJson{Error: "Unauthorized"})
		return "", "", nil
	}

	player := GetPlayerByToken(token, uuid)

	return token, uuid, player
}

func IsAuthorized(c *gin.Context) bool {
	token, uuid, _ := GetAuthorization(c)
	return token != "" && uuid != ""
}
