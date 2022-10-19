package endpoints

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func GetLobbyList(c *gin.Context) {
	if !database.IsAuthorized(c) {
		return
	}

	var lobbies []structs.LobbyInfo
	for i := 0; i < rand.Intn(10); i++ {
		lobbies = append(lobbies, structs.TempGenerateLobby())
	}

	c.JSON(200, structs.LobbyList{
		Lobbies: lobbies,
	})
}