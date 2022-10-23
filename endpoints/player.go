package endpoints

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"

	"github.com/gin-gonic/gin"
)

func GetPlayer(c *gin.Context) {
	if database.IsAuthorized(c) {
		return
	}
	
	var playerid = c.Param("playerid")
	if len(playerid) < 16 {
		c.JSON(400, structs.ErrorJson{
			Error: "invalid player id",
		})
		return
	}

	player := database.GetPlayerByID(playerid)
	if player == nil {
		c.JSON(404, structs.ErrorJson{
			Error: "player not found",
		})
		return
	}

	c.JSON(200, player)
}