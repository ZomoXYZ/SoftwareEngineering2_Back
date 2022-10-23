package endpoints

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"
	"edu/letu/wan/util"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	var token = util.GenerateToken()

	player := structs.GeneratePlayer()
	database.AddPlayer(token, player)

	c.JSON(200, structs.AuthorizationToken{
		Token: token,
	})
}

func CheckAuthorization(c *gin.Context) {
	var requestBody structs.AuthorizationToken

	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	player := database.GetPlayerByToken(requestBody.Token)
	//ERROR sql: no rows in result set

	if player == nil {
		c.AbortWithStatusJSON(401, structs.ErrorJson{
			Error: "Unauthorized",
		})
	} else {
		c.JSON(200, structs.AuthorizationToken{
			Token: requestBody.Token,
		})
	}
}