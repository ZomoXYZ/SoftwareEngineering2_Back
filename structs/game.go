package structs

type GamePlayer struct {
	Player Player
	Points int
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