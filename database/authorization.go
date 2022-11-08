package database

import (
	"edu/letu/wan/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthHeaders(r *http.Request) (string, string) {
	token := r.Header.Get("Authorization")
	uuid := r.Header.Get("UUID")

	return token, uuid
}

func GetAuthorizationPlayer(token string, uuid string) *structs.Player {
	exists := AuthorizationExists(token, uuid)
	if !exists {
		return nil
	}

	player := GetPlayerByToken(token, uuid)

	return player
}

func GetAuthorization(r *http.Request) (string, string, *structs.Player) {
	token, uuid := GetAuthHeaders(r)
	if token == "" || uuid == "" {
		return "", "", nil
	}

	player := GetAuthorizationPlayer(token, uuid)
	if player == nil {
		return "", "", nil
	}

	return token, uuid, player
}

func IsAuthorized(c *gin.Context) bool {
	token, uuid, _ := GetAuthorization(c.Request)
	fmt.Println(token, uuid)
	if token == "" || uuid == "" {
		c.AbortWithStatusJSON(401, structs.ErrorJson{Error: "Unauthorized"})
		return false
	}

	return true
}
