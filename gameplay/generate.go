package gameplay

import (
	"edu/letu/wan/structs"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func JsonLobbyWSFromGame(game *ActiveGame) string {
	var players = make([]*structs.Player, 0)
	for _, player := range game.Players {
		players = append(players, player.Player)
	}
    var lobbyWS = LobbyWS{
        ID: game.LobbyID,
        Code: game.LobbyCode,
        Host: game.Host.Player,
        Players: players,
    }
	lobbyJSON, err := json.Marshal(lobbyWS)
	if err != nil {
		fmt.Println("error converting lobby to json:", err)
		return ""
	}
	return string(lobbyJSON)
}

func GenerateActiveGame(lobby *structs.Lobby, host *structs.Player, hostConn *websocket.Conn) *ActiveGame {
	var game = ActiveGame{
		LobbyID: lobby.ID,
		LobbyCode: lobby.Code,
		Players: []*GamePlayer{},
		TurnState: TurnState{
			CurrentPlayer: 0,
			DidDraw: false,
			DidPlay: false,
			DidDiscard: false,
		},
		InLobby: true,

		join: make(chan *GamePlayer),
		leave: make(chan *GamePlayer),
		command: make(chan *PlayerCommandMessage),
	}
 
	var hostGamePlayer = GenerateGamePlayer(hostConn, host, &game)

	game.Host = hostGamePlayer

	return &game
}

func GenerateGamePlayer(conn *websocket.Conn, player *structs.Player, game *ActiveGame) *GamePlayer {
	var gamePlayer = GamePlayer{
		Player: player,
		Conn: conn,
		Points: 0,
		Game: game,

		send: make(chan CommandMessage),
	}

	return &gamePlayer
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
