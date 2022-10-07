package endpoints

import (
	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	c.JSON(200, gin.H{
		"token": "1234567890",
	})
}

type CheeckAuthorizationBody struct {
    Token string `json:"token" binding:"required"`
}

func CheckAuthorization(c *gin.Context) {
	var requestBody CheeckAuthorizationBody

	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	c.JSON(200, gin.H{
		"token": requestBody.Token,
		"valid": true,
	})
}