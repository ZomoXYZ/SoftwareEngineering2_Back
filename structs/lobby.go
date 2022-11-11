package structs

type LobbyInfo struct {
    ID string `json:"id" binding:"required"`
    Timestamp string `json:"timestamp" binding:"required"`
    Code string `json:"code" binding:"required"`
    Players int `json:"players" binding:"required"`
    Locked bool `json:"locked" binding:"required"`
}

type LobbyList struct {
	Lobbies []LobbyInfo `json:"lobbies" binding:"required"`
}

type LobbyWS struct {
    ID string `json:"id" binding:"required"`
    Code string `json:"code" binding:"required"`
    Host Player `json:"host" binding:"required"`
    Players []Player `json:"players" binding:"required"`
}

type Lobby struct {
    ID string
    Code string
    Password string
    Host Player
    Players []Player
    CreatedAt string
    HostJoined bool
    Started bool
}

func LobbyListFromLobbies(lobbies []Lobby) LobbyList {
    var lobbyList = LobbyList{ Lobbies: make([]LobbyInfo, 0) }
    for _, lobby := range lobbies {
        lobbyList.Lobbies = append(lobbyList.Lobbies, LobbyInfoFromLobby(lobby))
    }
    return lobbyList
}

func LobbyInfoFromLobby(lobby Lobby) LobbyInfo {
    return LobbyInfo{
        ID: lobby.ID,
        Timestamp: lobby.CreatedAt,
        Code: lobby.Code,
        Players: len(lobby.Players) + 1,
        Locked: lobby.Password != "",
    }
}

func LobbyWSFromLobby(lobby Lobby) LobbyWS {
    return LobbyWS{
        ID: lobby.ID,
        Code: lobby.Code,
        Host: lobby.Host,
        Players: lobby.Players,
    }
}
