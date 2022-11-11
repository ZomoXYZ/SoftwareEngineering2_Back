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
		sendMessage(conn, "unauthorized")
		conn.Close()
		return nil
	}

	var token = args[0]
	var uuid = args[1]

	// get player from database
	player := database.GetAuthorizationPlayer(token, uuid)
	if player == nil {
		sendMessage(conn, "unauthorized")
		conn.Close()
		return nil
	}

	sendMessage(conn, "authorized")

	return player
}

func connectLobby(conn *websocket.Conn, player *structs.Player) *structs.Lobby {
	command, args := readMessage(conn)

	// second command must be join lobby
	if command != "join" || len(args) < 1 {
		sendMessage(conn, "badcommand")
		conn.Close()
		return nil
	}

	var lobbyID = args[0]

	// get lobby from database
	lobby := database.GetLobby(lobbyID)
	if lobby == nil {
		sendMessage(conn, "badlobby")
		conn.Close()
		return nil
	}

	// make sure lobby can be joined
	if len(lobby.Players) >= 4 {
		sendMessage(conn, "lobbyfull")
		conn.Close()
		return nil
	}

	if lobby.Started {
		sendMessage(conn, "lobbyinprogress")
		conn.Close()
		return nil
	}

	database.JoinLobby(lobby.ID, *player)

	return lobby
}