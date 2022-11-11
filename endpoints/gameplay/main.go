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
}

type ConnCommandMessage struct {
	Cmd CommandMessage
	Conn *websocket.Conn
}

type PlayerCommandMessage struct {
	Cmd CommandMessage
	Player *GamePlayer
}

func Command(command string, args ...string) CommandMessage{
	return CommandMessage{
		Command: command,
		Args: args,
	}
}

func ConnCommand(conn *websocket.Conn, command string, args ...string) *ConnCommandMessage{
	return &ConnCommandMessage{
		Cmd: Command(command, args...),
		Conn: conn,
	}
}

func PlayerCommand(player *GamePlayer, conn *websocket.Conn, command string, args ...string) *PlayerCommandMessage{
	return &PlayerCommandMessage{
		Cmd: Command(command, args...),
		Player: player,
	}
}

func readMessage(conn *websocket.Conn) (*ConnCommandMessage, bool) {
	mt, message, err := conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			conn.Close()
			return nil, true
		}
		fmt.Println("read:", err)
		sendMessage(ConnCommand(conn, "error badmessage"))
		return nil, false
	}

	if mt != websocket.TextMessage {
		fmt.Println("weird websocket type:", mt)
		sendMessage(ConnCommand(conn, "error badmessage"))
		return nil, false
	}

	// parse command
	var split = strings.Split(string(message), " ")
	var command = split[0]
	var args = split[1:]

	fmt.Printf("recv command: %s\n     args: %s\n", command, strings.Join(args, " "))

	return ConnCommand(conn, command, args...), false
}

func sendMessage(command *ConnCommandMessage) {
	var fullMessage = command.Cmd.Command
	if len(command.Cmd.Args) > 0 {
		fullMessage += " " + strings.Join(command.Cmd.Args, " ")
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
