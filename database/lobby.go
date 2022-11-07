package database

import (
	"edu/letu/wan/structs"
	"sort"
	"time"
)

var Lobbies = make(map[string]structs.Lobby)

func AddLobby(host structs.Player) structs.Lobby {
	lobby := structs.GenerateLobby(host)
	// TODO check if user is already a host, give them that lobby if so
	Lobbies[lobby.ID] = lobby
	return lobby
}

func RemoveLobby(host structs.Player) {
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
		iCreatedAt, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", lobbyArray[i].CreatedAt)
		if err != nil {
			// TODO do something else ? this shouldn't ever trigger
			panic(err)
		}
		jCreatedAt, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", lobbyArray[j].CreatedAt)
		if err != nil {
			// TODO do something else ? this shouldn't ever trigger
			panic(err)
		}
		return iCreatedAt.After(jCreatedAt)
	})

	return lobbyArray
}

func UpdateLobbyPassword(lobbyid string, password string) {
	lobby := Lobbies[lobbyid]
	lobby.Password = password
	Lobbies[lobbyid] = lobby
}

func JoinLobby(lobbyid string, player structs.Player) {
	lobby := Lobbies[lobbyid]
	lobby.Players = append(lobby.Players, player.ID)
	Lobbies[lobbyid] = lobby
}

func LeaveLobby(lobbyid string, player structs.Player) {
	lobby := Lobbies[lobbyid]
	for i, id := range lobby.Players {
		if id == player.ID {
			lobby.Players = append(lobby.Players[:i], lobby.Players[i+1:]...)
		}
	}
}
