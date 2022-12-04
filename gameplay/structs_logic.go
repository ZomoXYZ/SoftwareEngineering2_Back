package gameplay

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"
	"edu/letu/wan/util"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// websocket rules
const (
	writeWait = 10 * time.Second
	pongWait = 30 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)

// mapping lobbyid to thread
type GameList map[string]*ActiveGame
var Games = make(GameList)

// connect player to lobby session
func joinLiveLobby(conn *websocket.Conn, player *structs.Player, lobby *structs.Lobby) {
	var gamePlayer *GamePlayer

	//find lobby game
	game, ok := Games[lobby.ID]
	if !ok {
		//create new game
		Games[lobby.ID] = GenerateActiveGame(lobby, player, conn)
		gamePlayer = Games[lobby.ID].Host
		go Games[lobby.ID].run()
	} else {
		//add player to game
		gamePlayer = GenerateGamePlayer(conn, player, game)
		game.Join <- gamePlayer
	}

	go gamePlayer.readPump()
	go gamePlayer.writePump()
}

// writer goroutine
func (p *GamePlayer) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		p.Conn.Close()
	}()
	for {
		select {
		case command, ok := <-p.Send:
			p.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				sendCloseMessage(p.Conn)
				return
			}
			
			sendMessage(ConnCommand(p.Conn, command.Command, command.Args...))
		case <-p.Close:
			sendCloseMessage(p.Conn)
			return
		case <-ticker.C:
			p.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := p.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// reader goroutine
func (p *GamePlayer) readPump() {
	defer func() {
		p.Conn.Close()
		p.Game.Leave <- p
	}()
	p.Conn.SetReadLimit(maxMessageSize)
	p.Conn.SetReadDeadline(time.Now().Add(pongWait))
	p.Conn.SetPongHandler(func(string) error { p.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		command, disconnected := readMessage(p.Conn)
		if disconnected {
			break
		}
		if command != nil {
			p.Game.Command <- PlayerCommand(p, command.Conn, command.Cmd.Command, command.Cmd.Args...)
		}
	}
}

// game session goroutine
func (game *ActiveGame) run() {
	//send data to host, who has just joined
	game.Host.Send <- Command("joined", JsonLobbyWSFromGame(game))

	//start listening
	for {
		select {
		// player joined
		case player := <-game.Join:
			if game.InLobby && len(game.Players) < 3 {
				// add player to game
				game.Players = append(game.Players, player)

				player.Send <- Command("joined", JsonLobbyWSFromGame(game))
				game.Broadcast(Command("playerupdate", JsonLobbyWSFromGame(game)), player)
			} else {
				// reject player
				player.Send <- Command("rejected")
				player.Close <- true
			}

		//player left
		case player := <-game.Leave:
			if player.Player.ID == game.Host.Player.ID {
				// host leaving
				// end game
				game.Close(true)
				return
			} else {
				// player leaving
				// remove player from game
				var index int
				for i, p := range game.Players {
					if p.Player.ID == player.Player.ID {
						index = i
						break
					}
				}
				game.Players = util.RemoveFromSlice(game.Players, index)

				game.Broadcast(Command("playerupdate", JsonLobbyWSFromGame(game)))

				// fix current player index
				if game.GameState.CurrentPlayer > index {
					game.GameState.CurrentPlayer--
				}

				game.NextTurn()
			}

		// command rom player
		case command := <-game.Command:
			fmt.Printf("Recv from: %s\n     command: %s\n     args: %s\n",
				command.Player.Player.ID, command.Cmd.Command, strings.Join(command.Cmd.Args, " "))

			if command.Player.Player.ID == game.Host.Player.ID {
				ran := RunHostCommand(game, command)
				if ran {
					continue
				}
			}

			ran := RunPlayerCommand(game, command)
			if ran {
				continue
			}
			command.Player.Send <- Command("error", "unknown command")
		}
	}
}

func (game *ActiveGame) Close(hostLeft bool) {
	fmt.Println("closing game")
	// don't send to host if they've already left
	if !hostLeft {
		// game.Host.Send <- Command("closed")
		game.Host.Close <- true
	}
	for _, player := range game.Players {
		// player.Send <- Command("closed")
		player.Close <- true
	}
	// remove game and lobby from lists
	database.RemoveLobby(*game.Host.Player)
	delete(Games, game.LobbyID)
}

func (game *ActiveGame) Broadcast(command CommandMessage, exclude ...*GamePlayer) {
	players := game.GetPlayers(exclude...)
	for _, player := range players {
		player.Send <- command
	}
}

func (game *ActiveGame) GetPlayers(exclude ...*GamePlayer) []*GamePlayer {
	var players []*GamePlayer
	players = append(players, game.Host)
	players = append(players, game.Players...)
	for _, ex := range exclude {
		for i, player := range players {
			if player.Player.ID == ex.Player.ID {
				players = util.RemoveFromSlice(players, i)
				break
			}
		}
	}
	return players
}

func (game *ActiveGame) ResetState(inLobby bool) {
	// game state
	game.InLobby = inLobby
	game.GameState = GameState{
		CurrentPlayer: 0,
		DiscardPile: structs.RandomCard(),
	}
	game.TurnState = TurnState{
		DidDraw: false,
		DidDiscard: false,
		DidPlay: false,
	}
	// player state
	for _, player := range game.GetPlayers() {
		player.Points = 0
		player.Cards = make([]structs.Card, 0)
		player.InGame = false
	}
	// lobby state
	database.GetLobby(game.LobbyID).Started = !inLobby
}

func (game *ActiveGame) NextTurn() {
	game.GameState.CurrentPlayer = (game.GameState.CurrentPlayer + 1) % len(game.GetPlayers())
	game.TurnState = TurnState{
		DidDraw: false,
		DidDiscard: false,
		DidPlay: false,
	}
	broadcastTurnState(game)
}