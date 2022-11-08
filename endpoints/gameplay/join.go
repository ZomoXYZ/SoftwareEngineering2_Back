package gameplay

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"

	"github.com/gorilla/websocket"
)


func connectPlayer(conn *websocket.Conn) *structs.Player {
	command, args := readMessage(conn)

	// first command must be authorization
	if command != "authorization" || len(args) < 2 {
		conn.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))
		conn.Close()
		return nil
	}

	var token = args[0]
	var uuid = args[1]

	// get player from database
	player := database.GetAuthorizationPlayer(token, uuid)
	if player == nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Unauthorized"))
		conn.Close()
		return nil
	}

	conn.WriteMessage(websocket.TextMessage, []byte("Ok"))

	return player
}

func connectLobby(conn *websocket.Conn, player *structs.Player) *structs.Lobby {
	command, args := readMessage(conn)

	// second command must be join lobby
	if command != "join" || len(args) < 1 {
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid Command"))
		conn.Close()
		return nil
	}

	var lobbyID = args[0]

	// get lobby from database
	lobby := database.GetLobby(lobbyID)
	if lobby == nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid Lobby"))
		conn.Close()
		return nil
	}

	// make sure lobby can be joined
	if len(lobby.Players) >= 4 {
		conn.WriteMessage(websocket.TextMessage, []byte("Lobby Full"))
		conn.Close()
		return nil
	}

	if lobby.Started {
		conn.WriteMessage(websocket.TextMessage, []byte("Lobby Started"))
		conn.Close()
		return nil
	}

	database.JoinLobby(lobby.ID, *player)

	conn.WriteMessage(websocket.TextMessage, []byte("Ok"))

	return lobby
}