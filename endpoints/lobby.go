package endpoints

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

func GetLobbyList(c *gin.Context) {
	var lobbies []LobbyInfo
	for i := 0; i < rand.Intn(10); i++ {
		lobbies = append(lobbies, TempGenerateLobby())
	}

	c.JSON(200, LobbyList{
		Lobbies: lobbies,
	})
}