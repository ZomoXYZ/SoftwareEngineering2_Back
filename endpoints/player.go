package endpoints

import (
	"github.com/gin-gonic/gin"
)

func GetPlayer(c *gin.Context) {
	var playerid = c.Param("playerid")

	if len(playerid) < 16 {
		c.JSON(400, ErrorJson{
			Error: "invalid player id",
		})
		return
	}

	//temp
	var player = TempGeneratePlayer()
	player.ID = playerid

	c.JSON(200, player)
}