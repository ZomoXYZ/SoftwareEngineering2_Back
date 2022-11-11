package gameplay

import (
	"edu/letu/wan/structs"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 30 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type LobbyWS struct {
    ID string `json:"id" binding:"required"`
    Code string `json:"code" binding:"required"`
    Host *structs.Player `json:"host" binding:"required"`
    Players []*structs.Player `json:"players" binding:"required"`
}


func JsonLobbyWSFromGame(game ActiveGame) string {
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

type GamePlayer struct {
	Player *structs.Player
	Conn *websocket.Conn
	Points int
	Game *ActiveGame

	send chan CommandMessage
}

type TurnState struct {
	CurrentPlayer int
	DidDraw bool
	DidPlay bool
	DidDiscard bool
}

type ActiveGame struct {
	LobbyID string
	LobbyCode string
	Host *GamePlayer
	Players []*GamePlayer
	TurnState TurnState
	InLobby bool

	join chan *GamePlayer
	leave chan *GamePlayer
	command chan *PlayerCommandMessage
	close chan bool
}

func GenerateActiveGame(lobby *structs.Lobby, host *structs.Player, hostConn *websocket.Conn) ActiveGame {
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
		close: make(chan bool),
	}
 
	var hostGamePlayer = GenerateGamePlayer(hostConn, host, &game)

	game.Host = hostGamePlayer

	return game
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

// mapping lobbyid to thread
type GameList map[string]ActiveGame
var Games = make(GameList)

func joinLiveLobby(conn *websocket.Conn, player *structs.Player, lobby *structs.Lobby) {

	var gamePlayer *GamePlayer

	//find lobby game
	game, ok := Games[lobby.ID]
	if !ok {
		fmt.Println("player starting game")
		//create new game
		Games[lobby.ID] = GenerateActiveGame(lobby, player, conn)
		gamePlayer = Games[lobby.ID].Host
		go Games[lobby.ID].run()
	} else {
		fmt.Println("player joining game")
		//add player to game
		gamePlayer = GenerateGamePlayer(conn, player, &game)
		game.join <- gamePlayer
	}

	go gamePlayer.writePump()
	go gamePlayer.readPump()

}

func (p *GamePlayer) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		p.Conn.Close()
	}()
	for {
		select {
		case command, ok := <-p.send:
			p.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				sendCloseMessage(p.Conn)
				return
			}
			
			sendMessage(ConnCommand(p.Conn, command.Command, command.Args...))
		case <-ticker.C:
			p.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := p.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (p *GamePlayer) readPump() {
	defer func() {
		p.Game.leave <- p
		p.Conn.Close()
	}()
	p.Conn.SetReadLimit(maxMessageSize)
	p.Conn.SetReadDeadline(time.Now().Add(pongWait))
	p.Conn.SetPongHandler(func(string) error { p.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		command, disconnected := readMessage(p.Conn)
		if disconnected {
			break;
		}
		if command != nil {
			p.Game.command <- PlayerCommand(p, command.Conn, command.Cmd.Command, command.Cmd.Args...)
		}
	}
}

// connect channels
func (game ActiveGame) run() {
	//send data to host, who has just joined
	game.Host.send <- Command("joined", JsonLobbyWSFromGame(game))

	//start listening
	for {
		select {
		// add player to game
		case player := <-game.join:
			if game.InLobby && len(game.Players) < 3 {
				game.Players = append(game.Players, player)

				player.send <- Command("joined", JsonLobbyWSFromGame(game))
				// TODO send message to all players that a new player has joined
			} else {
				// TODO reject player
			}

		case player := <-game.leave:
			fmt.Println("player leaving")
			if player.Player.ID == game.Host.Player.ID {
				// TODO end game
			} else {
				// remove player from game
				var index int
				for i, p := range game.Players {
					if p.Player.ID == player.Player.ID {
						index = i
						break
					}
				}
				game.Players = append(game.Players[:index], game.Players[index+1:]...)

				// TODO send message to all players that a player has left
			}

		case command := <-game.command:
			// TODO handle command from player
			fmt.Printf("Recv from: %s\n     command: %s\n     args: %s\n", command.Player.Player.ID, command.Cmd.Command, strings.Join(command.Cmd.Args, " "))

		// case close := <-game.close:
		// 	if close {
		// 		// TODO close all connections, close game, remove from Games
		// 	}
		}
	}
}

