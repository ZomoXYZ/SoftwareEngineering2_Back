package endpoints

import (
	"edu/letu/wan/util"
	"math/rand"
)

type AuthorizationToken struct {
	Token string `json:"token"`
}

type ErrorJson struct {
	Error string `json:"error" binding:"required"`
}

type LobbyInfo struct {
    ID string `json:"id" binding:"required"`
    Players int `json:"players" binding:"required"`
    Locked bool `json:"locked" binding:"required"`
}

type LobbyList struct {
	Lobbies []LobbyInfo `json:"lobbies" binding:"required"`
}

type PlayerName struct {
	Adjective int `json:"adjective" binding:"required"`
	Noun int `json:"noun" binding:"required"`
}

type PlayerInfo struct {
	ID string `json:"id" binding:"required"`
	Name PlayerName `json:"name" binding:"required"`
	Picture int `json:"picture" binding:"required"`
}

type SelfBody struct {
	Name *PlayerName `json:"name,omitempty"`
	Picture *int `json:"picture,omitempty"`
}

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