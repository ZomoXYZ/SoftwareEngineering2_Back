package endpoints

import (
	"edu/letu/wan/structs"

	"github.com/gin-gonic/gin"
)

func GetSelf(c *gin.Context) {
	//temp
	var player = structs.TempGeneratePlayer()

	c.JSON(200, player)
}

func SetSelf(c *gin.Context) {
	var requestBody structs.SelfRestBody

	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	if (requestBody.Name == nil && requestBody.Picture == nil) {
		c.JSON(400, structs.ErrorJson{
			Error: "invalid request body",
		})
		return
	}

	//temp
	var player = structs.TempGeneratePlayer()
	if (requestBody.Name != nil) {
		player.Name = *requestBody.Name
	}
	if (requestBody.Picture != nil) {
		player.Picture = *requestBody.Picture
	}

	c.JSON(200, player)
}
