package endpoints

import "github.com/gin-gonic/gin"

func Initialize(r *gin.Engine) {
	//generic/dev
	r.GET("/ping", Ping)

	//authorization
	r.GET("/authorization", Authorization)
	r.POST("/authorization", CheckAuthorization)

	//lobby
	r.GET("/lobbylist", GetLobbyList)

	//player
	r.GET("/player/:playerid", GetPlayer)

	//self
	r.GET("/self", GetSelf)
	r.POST("/self", SetSelf)

	//meta
	r.GET("/meta/names", MetaNames)
	r.GET("/meta/pictures", MetaPictures)
}