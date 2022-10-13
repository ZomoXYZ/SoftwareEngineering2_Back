package endpoints

import (
	"github.com/gin-gonic/gin"
)

func GetSelf(c *gin.Context) {
	//temp
	var player = TempGeneratePlayer()

	c.JSON(200, player)
}

func SetSelf(c *gin.Context) {
	var requestBody SelfBody

	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	if (requestBody.Name == nil && requestBody.Picture == nil) {
		c.JSON(400, ErrorJson{
			Error: "invalid request body",
		})
		return
	}

	//temp
	var player = TempGeneratePlayer()
	if (requestBody.Name != nil) {
		player.Name = *requestBody.Name
	}
	if (requestBody.Picture != nil) {
		player.Picture = *requestBody.Picture
	}

	c.JSON(200, player)
}
