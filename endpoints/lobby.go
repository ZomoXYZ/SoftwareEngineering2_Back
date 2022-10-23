package endpoints

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetLobbyList(c *gin.Context) {
	if database.IsAuthorized(c) {
		return
	}

	var lobbies = database.GetAvailableLobbies()
	var lobbyList = structs.LobbyListFromLobbies(lobbies)
	
	fmt.Println(lobbyList.Lobbies)
	fmt.Println(lobbyList)

	c.JSON(200, lobbyList)
}

func CreateLobby(c *gin.Context) {
	if database.IsAuthorized(c) {
		return
	}

	//TODO
}