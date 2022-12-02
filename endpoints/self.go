package endpoints

import (
	"edu/letu/wan/database"
	"edu/letu/wan/metauser"
	"edu/letu/wan/structs"
	"edu/letu/wan/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetSelf(c *gin.Context) {
	token, _, player := database.GetAuthorization(c.Request)
	if token == "" {
		return
	}

	c.JSON(200, player)
}

func SetSelf(c *gin.Context) {
	token, uuid, player := database.GetAuthorization(c.Request)
	if token == "" {
		return
	}
	
	var requestBody structs.RestBodySelf

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		return
	}

	names := metauser.GetMetaNames()
	avatars := metauser.GetMetaAvatars()

	if requestBody.Name == nil && requestBody.Picture == nil {
		c.JSON(400, structs.ErrorJson{
			Error: "invalid request body",
		})
		return
	}

	if requestBody.Name != nil {
		if requestBody.Name.Adjective != nil {
			adj := util.ValidateKeyFromMap(names.Adjectives, *requestBody.Name.Adjective)
			player.Name.Adjective = adj
		}
		if requestBody.Name.Noun != nil {
			noun := util.ValidateKeyFromMap(names.Nouns, *requestBody.Name.Noun)
			player.Name.Noun = noun
		}
	}
	if requestBody.Picture != nil {
		picture := util.ValidateKeyFromMap(avatars.Avatars, *requestBody.Picture)
		player.Picture = picture
	}

	database.UpdatePlayer(token, player)
	updatedPlayer := database.GetPlayerByToken(token, uuid)

	c.JSON(200, updatedPlayer)
}
