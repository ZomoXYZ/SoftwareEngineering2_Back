package gameplay

import (
	"edu/letu/wan/structs"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(req *http.Request) bool {
		// TODO check origin
		return true // temporarily prevent error from using localhost
	},
}

func readMessage(conn *websocket.Conn) (string, []string) {
	mt, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Error Reading Message"))
		return "", []string{}
	}

	if mt != websocket.TextMessage {
		log.Println("weird websocket type:", mt)
		conn.WriteMessage(websocket.TextMessage, []byte("Error Reading Message"))
		return "", []string{}
	}

	// parse command
	var split = strings.Split(string(message), " ")
	var command = split[0]
	var args = split[1:]

	log.Printf("recv command: %s\n     args: %s", command, strings.Join(args, " "))

	return command, args;
}

func playerInLobby(conn *websocket.Conn, player *structs.Player, lobby *structs.Lobby) {

	// TODO move this to another file
	// connect to game session
	// use structs.ActiveGame
	// each command will be a function and check against the state, return error if the state isn't correct
	// this for loop should listen for all 4 players at once, it'll choose which to listen to with structs.ActiveGame.Players[i].Conn

	for {
		command, args := readMessage(conn)

		// write message
		var err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("You sent command (%s) with args: %s", command, strings.Join(args, " "))))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	
}

func WSConnection(ginConn *gin.Context) {
	var w = ginConn.Writer
	var r = ginConn.Request

	// get as websocket command
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	//get player
	player := connectPlayer(conn)
	if player == nil {
		return
	}

	// join/get lobby
	lobby := connectLobby(conn, player)
	if lobby == nil {
		return
	}

	playerInLobby(conn, player, lobby)
}
