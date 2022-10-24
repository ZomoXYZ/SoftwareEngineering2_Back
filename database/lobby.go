package database

import (
	"edu/letu/wan/structs"
	"sort"
)

var Lobbies = make(map[string]structs.Lobby)

func AddLobby(host structs.PlayerInfo) structs.Lobby {
	lobby := structs.GenerateLobby(host)
	Lobbies[lobby.ID] = lobby
	return lobby
}

func RemoveLobby(host structs.PlayerInfo) {
	delete(Lobbies, host.ID)
}

func GetLobby(lobbyid string) *structs.Lobby {
	lobby, ok := Lobbies[lobbyid]
	if !ok {
		return nil
	}
	return &lobby
}

func GetAvailableLobbies() []structs.Lobby {
	//array of all lobbies
	lobbyArray := make([]structs.Lobby, 0, len(Lobbies))
	for  _, value := range Lobbies {
		lobbyArray = append(lobbyArray, value)
	}

	//sort by time created
	sort.Slice(lobbyArray, func(i, j int) bool {
		return lobbyArray[i].CreatedAt < lobbyArray[j].CreatedAt
	})

	return lobbyArray
}

func UpdateLobbyPassword(lobbyid string, password string) {
	lobby := Lobbies[lobbyid]
	lobby.Password = password
	Lobbies[lobbyid] = lobby
}

func JoinLobby(lobbyid string, player structs.PlayerInfo) {
	lobby := Lobbies[lobbyid]
	lobby.Players = append(lobby.Players, player.ID)
	Lobbies[lobbyid] = lobby
}

func LeaveLobby(lobbyid string, player structs.PlayerInfo) {
	lobby := Lobbies[lobbyid]
	for i, id := range lobby.Players {
		if id == player.ID {
			lobby.Players = append(lobby.Players[:i], lobby.Players[i+1:]...)
		}
	}
}
