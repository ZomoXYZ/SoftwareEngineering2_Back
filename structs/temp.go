package structs

import (
	"edu/letu/wan/util"
	"math/rand"
)

func TempGenerateLobby() LobbyInfo {
	var id = util.RandChars(6)
	var players = rand.Intn(3) + 1
	var locked = rand.Intn(2) == 1

	return LobbyInfo{
		ID: id,
		Players: players,
		Locked: locked,
	}
}

func TempGeneratePlayer() PlayerInfo {
	var playerid = util.RandNums(16)
	var nameAdj = rand.Intn(100)
	var nameNoun = rand.Intn(100)
	var picture = rand.Intn(100)

	return PlayerInfo{
		ID: playerid,
		Name: PlayerName{
			Adjective: nameAdj,
			Noun: nameNoun,
		},
		Picture: picture,
	}
}