package main

import (
	"edu/letu/wan/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	//generic/dev
	r.GET("/ping", endpoints.Ping)

	//authorization
	r.GET("/authorization", endpoints.Authorization)
	r.POST("/authorization", endpoints.CheckAuthorization)

	//lobby
	r.GET("/lobbylist", endpoints.GetLobbyList)

	//player
	r.GET("/player/:playerid", endpoints.GetPlayer)

	//self
	r.GET("/self", endpoints.GetSelf)
	r.POST("/self", endpoints.SetSelf)

	//meta
	r.GET("/meta/names", endpoints.MetaNames)
	r.GET("/meta/pictures", endpoints.MetaPictures)
        
	r.Run() // listen and serve on 0.0.0.0:8080
}