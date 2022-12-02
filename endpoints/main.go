package endpoints

import (
	"edu/letu/wan/gameplay"
	"flag"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

var (
	limit ratelimit.Limiter
	rps   = flag.Int("rps", 25, "request per second")
)

func Initialize(app *gin.Engine) {
	limit = ratelimit.New(*rps)
	
	app.Use(func(ctx *gin.Context) {
		limit.Take()
	})

	//generic/dev
	app.GET("/ping", Ping)
	app.GET("/pingauth", PingAuthorized)

	//authorization
	app.GET("/authorization", Authorization)
	app.POST("/authorization", CheckAuthorization)

	//lobby
	app.GET("/lobbylist", GetLobbyListLatest)
	app.GET("/lobbylist/:timestamp", GetLobbyListAfter)
	app.POST("/lobby", CreateLobby)
	app.GET("/lobby/:code", GetLobbyFromCode)

	//player
	app.GET("/player/:playerid", GetPlayer)

	//self
	app.GET("/self", GetSelf)
	app.POST("/self", SetSelf)

	//meta
	app.GET("/meta/names", MetaNames)
	app.GET("/meta/pictures", MetaPictures)
	app.GET("/meta/picture/:avatarID", MetaPictureServe)

	//websocket
	app.GET("/ws", gameplay.WSConnection)

	//teapot
	app.GET("/teapot", func (c *gin.Context) {
		c.AbortWithStatus(418)
	})
}