package gameplay

import (
	"edu/letu/wan/structs"

	"github.com/gorilla/websocket"
)

type CommandMessage struct {
	Command string
	Args []string
}

type ConnCommandMessage struct {
	Cmd CommandMessage
	Conn *websocket.Conn
}

type PlayerCommandMessage struct {
	Cmd CommandMessage
	Player *GamePlayer
}

type LobbyWS struct {
    ID string `json:"id" binding:"required"`
    Code string `json:"code" binding:"required"`
    Host string `json:"host" binding:"required"`
    Players []*structs.Player `json:"players" binding:"required"`
}

type GamePlayer struct {
	Player *structs.Player
	Conn *websocket.Conn
	Points int
	Cards []structs.Card
	Game *ActiveGame
	InGame bool

	Send chan CommandMessage
	Close chan bool
}

type TurnState struct {
	DidDraw bool
	DidDiscard bool
	DidPlay bool
}
type GameState struct {
	EveryoneIn bool
	CurrentPlayer int
	DiscardPile structs.Card //only need to store top card
}

type GameSettings struct {
	PointsToWin int
}

type ActiveGame struct {
	LobbyID string
	LobbyCode string
	Host *GamePlayer
	Players []*GamePlayer
	TurnState TurnState
	GameState GameState
	InLobby bool

	Settings GameSettings

	Join chan *GamePlayer
	Leave chan *GamePlayer
	Command chan *PlayerCommandMessage
}