package endpoints

import (
	"edu/letu/wan/util"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	//temp
	var token = util.RandAll(32)

	c.JSON(200, AuthorizationToken{
		Token: token,
	})
}

func CheckAuthorization(c *gin.Context) {
	var requestBody AuthorizationToken

	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	c.JSON(200, gin.H{
		"token": requestBody.Token,
		"valid": true,
	})
}