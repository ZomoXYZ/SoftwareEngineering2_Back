package endpoints

import (
	"edu/letu/wan/structs"
	"edu/letu/wan/util"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	//temp
	var token = util.RandAll(32)

	// this will generate a randomized user, the client may change it immediately after with /self

	c.JSON(200, structs.AuthorizationToken{
		Token: token,
	})
}

func CheckAuthorization(c *gin.Context) {
	var requestBody structs.AuthorizationToken

	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	c.JSON(200, gin.H{
		"token": requestBody.Token,
		"valid": true,
	})
}