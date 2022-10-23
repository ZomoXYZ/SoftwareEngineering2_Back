package structs

type LobbyInfo struct {
    ID string `json:"id" binding:"required"`
    Players int `json:"players" binding:"required"`
    Locked bool `json:"locked" binding:"required"`
}

type LobbyList struct {
	Lobbies []LobbyInfo `json:"lobbies" binding:"required"`
}

type Lobby struct {
    ID string
    Code string
    Password string
    Host string
    Players []string
    CreatedAt string
}

func LobbyListFromLobbies(lobbies []Lobby) LobbyList {
    var lobbyList = LobbyList{ Lobbies: make([]LobbyInfo, 0) }
    for _, lobby := range lobbies {
        lobbyList.Lobbies = append(lobbyList.Lobbies, LobbyInfo{
            ID: lobby.ID,
            Players: len(lobby.Players) + 1,
            Locked: lobby.Password != "",
        })
    }
    return lobbyList
}
