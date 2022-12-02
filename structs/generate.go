package structs

import (
	"edu/letu/wan/metauser"
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

	names := metauser.GetMetaNames()

	var nameAdj = util.RandomKeyFromMap(names.Adjectives)
	var nameNoun = util.RandomKeyFromMap(names.Nouns)
	var picture = rand.Intn(100) // TODO pictures

	return Player{
		ID: playerid.String(),
		Name: PlayerName{
			Adjective: nameAdj,
			Noun: nameNoun,
		},
		Picture: picture,
	}
}

func GenerateLobby(host *Player) *Lobby {
	node, err := snowflake.NewNode(nodeIndex)
	nodeIndex++;
	if err != nil {
		panic(err)
	}

	var lobbyid = node.Generate()

	return &Lobby{
		ID: lobbyid.String(),
		Code: util.LobbyCode(),
		Password: "",
		Host: host,
		Players: []*Player{},
		CreatedAt: time.Now().UTC().String(),
		HostJoined: false,
		Started: false,
	}
}
