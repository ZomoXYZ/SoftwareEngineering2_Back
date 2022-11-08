package structs

import "github.com/gorilla/websocket"

type PlayerRole int

const (
	RoleHost PlayerRole = iota
	RolePlayer
)

type GamePlayer struct {
	Player Player
	Conn *websocket.Conn
	Points int
	Role PlayerRole
}

type TurnState struct {
	CurrentPlayer int
	DidDraw bool
	DidPlay bool
	DidDiscard bool
}

type ActiveGame struct {
	LobbyID string
	Players []GamePlayer
	TurnState TurnState
}