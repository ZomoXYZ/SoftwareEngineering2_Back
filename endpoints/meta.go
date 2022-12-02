package endpoints

import (
	"edu/letu/wan/database"
	"edu/letu/wan/metauser"
	"edu/letu/wan/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MetaNames(c *gin.Context) {
	if !database.IsAuthorized(c) {
		return
	}

	payload := metauser.GetMetaNames()
	c.JSON(200, payload)
}

func MetaPictures(c *gin.Context) {
	if !database.IsAuthorized(c) {
		return
	}

	payload := metauser.GetMetaAvatarKeys()
	c.JSON(200, payload)
}

func MetaPictureServe(c *gin.Context) {
	if !database.IsAuthorized(c) {
		return
	}

	avatars := metauser.GetMetaAvatars().Avatars
	avatarID := c.Param("avatarID")

	avatarInt, err := strconv.Atoi(avatarID)
	if err != nil {
		c.JSON(400, structs.ErrorJson{
			Error: "invalid avatar id",
		})
		return
	}

	c.File("./resources/avatars/" + avatars[avatarInt])
}