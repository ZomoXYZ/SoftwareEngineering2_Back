package gameplay

import (
	"edu/letu/wan/structs"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func JsonLobbyWSFromGame(game *ActiveGame) string {
	gamePlayers := game.GetPlayers()
	players := []*structs.Player{}
	for _, gamePlayer := range gamePlayers {
		players = append(players, gamePlayer.Player)
	}
    var lobbyWS = LobbyWS{
        ID: game.LobbyID,
        Code: game.LobbyCode,
        Host: game.Host.Player.ID,
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
			DidDraw: false,
			DidPlay: false,
			DidDiscard: false,
		},
		GameState: GameState{
			CurrentPlayer: 0,
			DiscardPile: structs.RandomCard(),
		},
		InLobby: true,

		Settings: GameSettings{
			PointsToWin: 17,
		},

		Join: make(chan *GamePlayer),
		Leave: make(chan *GamePlayer),
		Command: make(chan *PlayerCommandMessage),
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

		Send: make(chan CommandMessage),
		Close: make(chan bool),
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
