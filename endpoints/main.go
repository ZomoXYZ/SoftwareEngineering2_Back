package endpoints

import (
	"edu/letu/wan/endpoints/gameplay"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	//generic/dev
	r.GET("/ping", Ping)
	r.GET("/pingauth", PingAuthorized)

	//authorization
	r.GET("/authorization", Authorization)
	r.POST("/authorization", CheckAuthorization)

	//lobby
	r.GET("/lobbylist", GetLobbyListLatest)
	r.GET("/lobbylist/:timestamp", GetLobbyListAfter)
	r.POST("/lobby", CreateLobby)
	r.GET("/lobby/:code", GetLobbyFromCode)

	// TODO remove these
	//temp
	r.GET("/createlobbies", TempCreateLobbies)
	r.GET("/deletelobbies", TempDeleteLobbies)

	//player
	r.GET("/player/:playerid", GetPlayer)

	//self
	r.GET("/self", GetSelf)
	r.POST("/self", SetSelf)

	//meta
	r.GET("/meta/names", MetaNames)
	r.GET("/meta/pictures", MetaPictures)

	//websocket
	r.GET("/ws", gameplay.WSConnection)

	//teapot
	r.GET("/teapot", func (c *gin.Context) {
		c.AbortWithStatus(418)
	})
}