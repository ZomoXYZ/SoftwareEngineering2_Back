package database

import (
	"edu/letu/wan/structs"
	"testing"
)

/*
AddLobby
RemoveLobby
GetLobby
GetAvailableLobbies
UpdateLobbyPassword
JoinLobby
LeaveLobby
*/

func TestAddLobby(t *testing.T) {
	// clear lobbies
	Lobbies = make(map[string]structs.Lobby)

	// create a player
	host := structs.PlayerInfo{
		ID: "123",
		Name: structs.PlayerName{
			Adjective: 1,
			Noun: 1,
		},
		Picture: 1,
	}

	lobby := AddLobby(host)

	if lobby.Host != host.ID {
		t.Errorf("Lobby host was not set correctly")
	}

}

func TestGetLobby(t *testing.T) {
	// clear lobbies
	Lobbies = make(map[string]structs.Lobby)

	// create a player
	host := structs.PlayerInfo{
		ID: "123",
		Name: structs.PlayerName{
			Adjective: 1,
			Noun: 1,
		},
		Picture: 1,
	}

	lobby := AddLobby(host)

	if lobby.Host != host.ID {
		t.Errorf("Lobby host was not set correctly")
	}

}
