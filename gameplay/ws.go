package gameplay

import (
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

func readMessage(conn *websocket.Conn) (*ConnCommandMessage, bool) {
	mt, message, err := conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err) {
			conn.Close()
			return nil, true
		}
		sendMessage(ConnCommand(conn, "error badmessage"))
		return nil, false
	}

	if mt == websocket.CloseMessage {
		conn.Close()
		return nil, true
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
		fmt.Printf("MSG WRITE ERROR: %s, attempting to send %s\n", err, command.Cmd.Command)
	}
}

func sendCloseMessage(conn *websocket.Conn) {
	err := conn.WriteMessage(websocket.CloseMessage, []byte{})
	if err != nil {
		fmt.Println("CLOSE MSG WRITE ERROR:", err)
	}
}