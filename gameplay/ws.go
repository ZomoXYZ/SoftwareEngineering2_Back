package gameplay

import (
	"fmt"
	"io"
	"strings"

	"github.com/gorilla/websocket"
)

func readMessage(conn *websocket.Conn) (*ConnCommandMessage, bool) {

	//get reader
	var r io.Reader
	messageType, r, err := conn.NextReader()
	if err != nil {
		// error in connection, player is offline
		return nil, true
	}

	//read message
	message, err := io.ReadAll(r)
	if err != nil {
		sendMessage(ConnCommand(conn, "error badmessage"))
		return nil, false
	}

	if messageType == websocket.CloseMessage {
		conn.Close()
		return nil, true
	}

	if messageType != websocket.TextMessage {
		fmt.Println("weird websocket type:", messageType)
		sendMessage(ConnCommand(conn, "error badmessage"))
		return nil, false
	}

	// parse command
	var split = strings.Split(string(message), " ")
	var command = split[0]
	var args = split[1:]

	// fmt.Printf("recv command: %s\n     args: %s\n", command, strings.Join(args, " "))

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