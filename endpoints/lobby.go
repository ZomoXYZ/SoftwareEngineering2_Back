package endpoints

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const LobbiesPerPage = 20

func GetLobbyListLatest(c *gin.Context) {
	if !database.IsAuthorized(c) {
		return
	}

	var lobbies = database.GetAvailableLobbies()

	var endIndex = LobbiesPerPage
	if len(lobbies) < LobbiesPerPage {
		endIndex = len(lobbies)
	}

	lobbies = lobbies[:endIndex]
	var lobbyList = structs.LobbyListFromLobbies(lobbies)

	c.JSON(200, lobbyList)
}

func GetLobbyListAfter(c *gin.Context) {
	if !database.IsAuthorized(c) {
		return
	}

	//2022-11-04 15:13:30.024317 +0000 UTC
	timestamp, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", c.Param("timestamp"))
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	// find the next oldest lobby than the timestamp
	var lobbies = database.GetAvailableLobbies()
	var startIndex = -1
	for i, lobby := range lobbies {
		createdAt, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", lobby.CreatedAt)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		if createdAt.Before(timestamp) {
			startIndex = i
			break
		}
	}

	// if no lobbies were created after the timestamp, return empty list
	if startIndex == -1 {
		c.JSON(200, structs.LobbyListFromLobbies([]*structs.Lobby{}))
		return
	}

	// if there are less lobbies than the page size, return all of them
	var endIndex = startIndex + LobbiesPerPage
	if len(lobbies) < endIndex {
		endIndex = len(lobbies)
	}

	//slice lobby array and return
	lobbies = lobbies[startIndex:endIndex]
	var lobbyList = structs.LobbyListFromLobbies(lobbies)

	c.JSON(200, lobbyList)
}

func CreateLobby(c *gin.Context) {
	_, _, player := database.GetAuthorization(c.Request)
	if player == nil {
		return
	}

	lobby := database.AddLobby(*player)

	c.JSON(200, structs.LobbyInfoFromLobby(*lobby))
}

func GetLobbyFromCode(c *gin.Context) {
	_, _, player := database.GetAuthorization(c.Request)
	if player == nil {
		return
	}

	fmt.Println("getting lobby from code")

	code := c.Param("code")
	lobby := database.GetLobbyFromCode(code)
	if lobby == nil {
		c.AbortWithStatus(404)
		fmt.Println("no lobby")
		return
	}
	fmt.Println("got lobby from code")

	c.JSON(200, structs.LobbyInfoFromLobby(*lobby))
}

func TempCreateLobbies(c *gin.Context) {
	player := structs.GeneratePlayer()

	for i := 0; i < 100; i++ {
		database.AddLobby(player)
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func TempDeleteLobbies(c *gin.Context) {
	database.Lobbies = make(map[string]*structs.Lobby)

	c.JSON(200, gin.H{
		"message": "success",
	})
}