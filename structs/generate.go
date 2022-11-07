package structs

import (
	"edu/letu/wan/util"
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

var nodeIndex int64 = 0

func GeneratePlayer() Player {
	node, err := snowflake.NewNode(nodeIndex)
	nodeIndex++;
	if err != nil {
		panic(err)
	}

	var playerid = node.Generate()
	// 100 is temporary
	var nameAdj = rand.Intn(100)
	var nameNoun = rand.Intn(100)
	var picture = rand.Intn(100)

	return Player{
		ID: playerid.String(),
		Name: PlayerName{
			Adjective: nameAdj,
			Noun: nameNoun,
		},
		Picture: picture,
	}
}

func GenerateLobby(host Player) Lobby {
	node, err := snowflake.NewNode(nodeIndex)
	nodeIndex++;
	if err != nil {
		panic(err)
	}

	var lobbyid = node.Generate()

	return Lobby{
		ID: lobbyid.String(),
		Code: util.LobbyCode(),
		Password: "",
		Host: host.ID,
		Players: []string{},
		CreatedAt: time.Now().UTC().String(),
	}
}
