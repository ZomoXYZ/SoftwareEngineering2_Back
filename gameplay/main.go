package gameplay

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(req *http.Request) bool {
		// TODO check origin
		return true // temporarily prevent error from using localhost
	},
}

func WSConnection(ginConn *gin.Context) {
	var w = ginConn.Writer
	var r = ginConn.Request

	// get as websocket command
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}

	//get player
	player := connectPlayer(conn)
	if player == nil { // connection already closed if nil
		return
	}

	// join/get lobby
	lobby := connectLobby(conn, player)
	if lobby == nil { // connection already closed if nil
		return
	}

	joinLiveLobby(conn, player, lobby)
}
