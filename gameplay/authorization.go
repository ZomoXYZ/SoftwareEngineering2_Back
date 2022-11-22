package gameplay

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"

	"github.com/gorilla/websocket"
)


func connectPlayer(conn *websocket.Conn) *structs.Player {
	command, disconnected := readMessage(conn)
	if disconnected {
		return nil
	}

	// first command must be authorization
	if command.Cmd.Command != "authorization" || len(command.Cmd.Args) < 2 {
		sendMessage(ConnCommand(conn, "unauthorized"))
		conn.Close()
		return nil
	}

	var (
		token = command.Cmd.Args[0]
		uuid = command.Cmd.Args[1]
	)

	// get player from database
	player := database.GetAuthorizationPlayer(token, uuid)
	if player == nil {
		sendMessage(ConnCommand(conn, "unauthorized"))
		conn.Close()
		return nil
	}

	sendMessage(ConnCommand(conn, "authorized"))

	return player
}

func connectLobby(conn *websocket.Conn, player *structs.Player) *structs.Lobby {
	
	command, disconnected := readMessage(conn)
	if disconnected {
		return nil
	}

	// second command must be join lobby
	if command.Cmd.Command != "join" || len(command.Cmd.Args) < 1 {
		sendMessage(ConnCommand(conn, "badcommand"))
		conn.Close()
		return nil
	}

	var lobbyID = command.Cmd.Args[0]

	// get lobby from database
	lobby := database.GetLobby(lobbyID)
	if lobby == nil {
		sendMessage(ConnCommand(conn, "badlobby"))
		conn.Close()
		return nil
	}

	// check password
	if lobby.Password != "" {
		// second arg will be password
		if len(command.Cmd.Args) < 2 || command.Cmd.Args[1] != lobby.Password {
			sendMessage(ConnCommand(conn, "badpassword"))
			conn.Close()
			return nil
		}
	}

	// make sure lobby can be joined
	if len(lobby.Players) >= 4 {
		sendMessage(ConnCommand(conn, "lobbyfull"))
		conn.Close()
		return nil
	}

	if lobby.Started {
		sendMessage(ConnCommand(conn, "lobbyinprogress"))
		conn.Close()
		return nil
	}
	
	//check if lobby is waiting for host, check if player is host
	if !lobby.HostJoined {
		//check if player is host
		if player.ID != lobby.Host.ID {
			sendMessage(ConnCommand(conn, "badlobby"))
			conn.Close()
			return nil
		}
		joinedLobby := database.HostJoinLobby(lobby.ID)
		if joinedLobby == nil {
			sendMessage(ConnCommand(conn, "badlobby"))
			conn.Close()
			return nil
		}
		return joinedLobby
	}

	// add player to lobby
	joinedLobby := database.JoinLobby(lobby.ID, *player)
	if joinedLobby == nil {
		sendMessage(ConnCommand(conn, "badlobby"))
		conn.Close()
		return nil
	}
	return joinedLobby
}