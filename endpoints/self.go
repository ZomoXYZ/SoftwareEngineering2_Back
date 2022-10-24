package endpoints

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"

	"github.com/gin-gonic/gin"
)

func GetSelf(c *gin.Context) {
	token, _, player := database.GetAuthorization(c)
	if token == "" {
		return
	}

	c.JSON(200, player)
}

func SetSelf(c *gin.Context) {
	token, _, player := database.GetAuthorization(c)
	if token == "" {
		return
	}
	
	var requestBody structs.RestBodySelf

	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	if (requestBody.Name == nil && requestBody.Picture == nil) {
		c.JSON(400, structs.ErrorJson{
			Error: "invalid request body",
		})
		return
	}

	if (requestBody.Name != nil) {
		player.Name = *requestBody.Name
	}
	if (requestBody.Picture != nil) {
		player.Picture = *requestBody.Picture
	}

	database.UpdatePlayer(token, player)
	updatedPlayer := database.GetPlayerByToken(token)

	c.JSON(200, updatedPlayer)
}
