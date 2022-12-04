package database

import (
	"edu/letu/wan/structs"
	"edu/letu/wan/util"
	"fmt"
	"sort"
	"time"
)

var Lobbies = make(map[string]*structs.Lobby)

func AddLobby(host structs.Player) *structs.Lobby {
	lobby := GetLobbyByHost(host)
	if lobby != nil {
		return lobby
	}
	lobby = structs.GenerateLobby(&host)
	Lobbies[lobby.ID] = lobby
	return lobby
}

func RemoveLobby(host structs.Player) {
	fmt.Println("starting to remove lobby for host", host.ID)
	lobby := GetLobbyByHost(host)
	if lobby != nil {
		fmt.Println("removing lobby", lobby.Code)
		delete(Lobbies, lobby.ID)
	} else {
		fmt.Println("removing lobby NONE TO REMOVE")
	}
}

func GetLobby(lobbyid string) *structs.Lobby {
	lobby, ok := Lobbies[lobbyid]
	if !ok {
		return nil
	}
	return lobby
}

func GetLobbyByHost(host structs.Player) *structs.Lobby {
	for _, value := range Lobbies {
		if value.Host.ID == host.ID {
			return value
		}
	}
	return nil
}

func GetLobbyFromCode(code string) *structs.Lobby {
	for _, value := range Lobbies {
		fmt.Println(code, value.Code)
		if value.Code == code {
			return value
		}
	}
	return nil
}

func GetAvailableLobbies() []*structs.Lobby {
	//array of all lobbies
	lobbyArray := make([]*structs.Lobby, 0, len(Lobbies))
	for  _, value := range Lobbies {
		
		// if lobby's host has joined, is not started, and is not full, then add to array
		if value.HostJoined && !value.Started && len(value.Players) < 4 {
			lobbyArray = append(lobbyArray, value)
		}
	}

	//sort by time created
	sort.Slice(lobbyArray, func(i, j int) bool {
		iCreatedAt, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", lobbyArray[i].CreatedAt)
		if err != nil {
			// return as if i is older
			return false
		}
		jCreatedAt, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", lobbyArray[j].CreatedAt)
		if err != nil {
			// return as if j is older
			return true
		}
		return iCreatedAt.After(jCreatedAt)
	})



	return lobbyArray
}

func UpdateLobbyPassword(lobbyid string, password string) {
	if lobby, ok := Lobbies[lobbyid]; ok {
		lobby.Password = password
		Lobbies[lobbyid] = lobby
	}
}

func JoinLobby(lobbyid string, player structs.Player) *structs.Lobby {
	if lobby, ok := Lobbies[lobbyid]; ok {
		if len(lobby.Players) < 4 {
			lobby.Players = append(lobby.Players, &player)
			Lobbies[lobbyid] = lobby
			return lobby
		}
	}
	return nil
}

func HostJoinLobby(lobbyid string) *structs.Lobby {
	if lobby, ok := Lobbies[lobbyid]; ok {
		lobby.HostJoined = true
		Lobbies[lobbyid] = lobby
		return lobby
	}
	return nil
}

func LeaveLobby(lobbyid string, player structs.Player) {
	if lobby, ok := Lobbies[lobbyid]; ok {
		for i, p := range lobby.Players {
			if p.ID == player.ID {
				lobby.Players = util.RemoveFromSlice(lobby.Players, i)
				Lobbies[lobbyid] = lobby
			}
		}
	}
}
