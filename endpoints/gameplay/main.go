package gameplay

import (
	"edu/letu/wan/database"
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

func readMessage(conn *websocket.Conn) (string, []string, int) {
	mt, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		conn.WriteMessage(mt, []byte("Error Reading Message"))
		return "", []string{}, mt
	}

	// parse command
	var split = strings.Split(string(message), " ")
	var command = split[0]
	var args = split[1:]

	log.Printf("recv command: %s\n     args: %s", command, strings.Join(args, " "))

	return command, args, mt;
}

func connectPlayer(conn *websocket.Conn) *structs.Player {
	command, args, mt := readMessage(conn)

	// first command must be authorization
	if command != "authorization" || len(args) < 2 {
		conn.WriteMessage(mt, []byte("Unauthorized"))
		conn.Close()
		return nil
	}

	var token = args[0]
	var uuid = args[1]

	// get player from database
	player := database.GetAuthorizationPlayer(token, uuid)
	if player == nil {
		conn.WriteMessage(mt, []byte("Unauthorized"))
		conn.Close()
		return nil
	}

	return player
}

func connectLobby(conn *websocket.Conn, player *structs.Player) *structs.Lobby {
	command, args, mt := readMessage(conn)

	// second command must be join lobby
	if command != "join" || len(args) < 1 {
		conn.WriteMessage(mt, []byte("Invalid Command"))
		conn.Close()
		return nil
	}

	var lobbyID = args[0]

	// get lobby from database
	lobby := database.GetLobby(lobbyID)
	if lobby == nil {
		conn.WriteMessage(mt, []byte("Invalid Lobby"))
		conn.Close()
		return nil
	}

	// make sure lobby can be joined
	if len(lobby.Players) >= 4 {
		conn.WriteMessage(mt, []byte("Lobby Full"))
		conn.Close()
		return nil
	}

	// TODO either we let people join a game in progress and remove the following code, or don't let people join mid game and add a second flag to the lobby of Starting, Starting would let players join and Started would say the game is actually in progress

	// if lobby.Started {
	// 	log.Println("read:", err)
	// 	c.WriteMessage(mt, []byte("Lobby Started"))
	// 	c.Close()
	// 	return nil
	// }

	database.JoinLobby(lobby.ID, *player)

	return lobby
}

func playerInLobby(conn *websocket.Conn, player *structs.Player, lobby *structs.Lobby) {

	// TODO move this to another file
	// connect to game session
	// use structs.ActiveGame
	// each command will be a function and check against the state, return error if the state isn't correct
	// this for loop should listen for all 4 players at once, it'll choose which to listen to with structs.ActiveGame.Players[i].Conn

	for {
		command, args, mt := readMessage(conn)

		// write message
		var err = conn.WriteMessage(mt, []byte(fmt.Sprintf("You sent command (%s) with args: %s", command, strings.Join(args, " "))))
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
