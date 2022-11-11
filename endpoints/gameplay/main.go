package gameplay

import (
	"fmt"
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

type CommandMessage struct {
	Command string
	Args []string
	Conn *websocket.Conn
}

func Command(conn *websocket.Conn, command string, args ...string) *CommandMessage{
	return &CommandMessage{
		Command: command,
		Args: args,
		Conn: conn,
	}
}

func readMessage(conn *websocket.Conn) *CommandMessage {
	mt, message, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("read:", err)
		sendMessage(Command(conn, "error badmessage"))
		return nil
	}

	if mt != websocket.TextMessage {
		fmt.Println("weird websocket type:", mt)
		sendMessage(Command(conn, "error badmessage"))
		return nil
	}

	// parse command
	var split = strings.Split(string(message), " ")
	var command = split[0]
	var args = split[1:]

	fmt.Printf("recv command: %s\n     args: %s", command, strings.Join(args, " "))

	return &CommandMessage{
		Command: command,
		Args: args,
		Conn: conn,
	};
}

func sendMessage(command *CommandMessage) {
	var fullMessage = command.Command
	if len(command.Args) > 0 {
		fullMessage += " " + strings.Join(command.Args, " ")
	}
	err := command.Conn.WriteMessage(websocket.TextMessage, []byte(fullMessage))
	if err != nil {
		fmt.Println("write:", err)
	}
}

func sendCloseMessage(conn *websocket.Conn) {
	err := conn.WriteMessage(websocket.CloseMessage, []byte{})
	if err != nil {
		fmt.Println("write:", err)
	}
}

// func joinLiveLobby(conn *websocket.Conn, player *structs.Player, lobby *structs.Lobby) {
// 	//player is in lobby, host is known to be connected
// 	// if lobby already has a thread, look up the thread and send the player to it
// 	// if lobby does not have a thread, create a new thread and send the player to it

// 	// TODO move this to another file
// 	// connect to game session
// 	// use structs.ActiveGame
// 	// each command will be a function and check against the state, return error if the state isn't correct
// 	// this for loop should listen for all 4 players at once, it'll choose which to listen to with structs.ActiveGame.Players[i].Conn

// 	for {
// 		command := readMessage(conn)
// 		sendMessage(command)
// 	}
	
// }

func WSConnection(ginConn *gin.Context) {
	var w = ginConn.Writer
	var r = ginConn.Request

	// get as websocket command
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("upgrade:", err)
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

	// send lobby data to player
	// lobbyWS := structs.LobbyWSFromLobby(*lobby)
	// lobbyJSON, err := json.Marshal(lobbyWS)
	// if err != nil {
	// 	fmt.Println("error converting lobby to json:", err)
	// 	return
	// }
	// // var message = fmt.Sprintf("joined %s", string(lobbyJSON))
	// // sendMessage(Command(conn, message))

	joinLiveLobby(conn, player, lobby)
}
