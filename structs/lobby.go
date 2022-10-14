package structs

type LobbyInfo struct {
    ID string `json:"id" binding:"required"`
    Players int `json:"players" binding:"required"`
    Locked bool `json:"locked" binding:"required"`
}

type LobbyList struct {
	Lobbies []LobbyInfo `json:"lobbies" binding:"required"`
}