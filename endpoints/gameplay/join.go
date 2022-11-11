package gameplay

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"

	"github.com/gorilla/websocket"
)


func connectPlayer(conn *websocket.Conn) *structs.Player {
	command := readMessage(conn)

	// first command must be authorization
	if command.Command != "authorization" || len(command.Args) < 2 {
		sendMessage(Command(conn, "unauthorized"))
		conn.Close()
		return nil
	}

	var token = command.Args[0]
	var uuid = command.Args[1]

	// get player from database
	player := database.GetAuthorizationPlayer(token, uuid)
	if player == nil {
		sendMessage(Command(conn, "unauthorized"))
		conn.Close()
		return nil
	}

	sendMessage(Command(conn, "authorized"))

	return player
}

func connectLobby(conn *websocket.Conn, player *structs.Player) *structs.Lobby {
	command := readMessage(conn)

	// second command must be join lobby
	if command.Command != "join" || len(command.Args) < 1 {
		sendMessage(Command(conn, "badcommand"))
		conn.Close()
		return nil
	}

	var lobbyID = command.Args[0]

	// get lobby from database
	lobby := database.GetLobby(lobbyID)
	if lobby == nil {
		sendMessage(Command(conn, "badlobby"))
		conn.Close()
		return nil
	}

	// make sure lobby can be joined
	if len(lobby.Players) >= 4 {
		sendMessage(Command(conn, "lobbyfull"))
		conn.Close()
		return nil
	}

	if lobby.Started {
		sendMessage(Command(conn, "lobbyinprogress"))
		conn.Close()
		return nil
	}
	
	//check if lobby is waiting for host, check if player is host
	if !lobby.HostJoined {
		//check if player is host
		if player.ID != lobby.Host.ID {
			sendMessage(Command(conn, "badlobby"))
			conn.Close()
			return nil
		}
		joinedLobby := database.HostJoinLobby(lobby.ID)
		if joinedLobby == nil {
			sendMessage(Command(conn, "badlobby"))
			conn.Close()
			return nil
		}
		return joinedLobby
	}

	// add player to lobby
	joinedLobby := database.JoinLobby(lobby.ID, *player)
	if joinedLobby == nil {
		sendMessage(Command(conn, "badlobby"))
		conn.Close()
		return nil
	}
	return joinedLobby
}