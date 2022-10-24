package database

import (
	"edu/letu/wan/structs"
	"reflect"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	ClearPlayerTable()

	var playerInfo = structs.PlayerInfo{
		ID: "1",
		Name: structs.PlayerName{
			Adjective: 1,
			Noun:      2,
		},
		Picture: 3,
	}

	//add one player
	AddPlayer("TOKEN", "UUID", playerInfo)

	//initialize test variables
	var foundPlayer *structs.PlayerInfo

	//check if getters work
	foundPlayer = GetPlayerByID("1")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo)) {
		t.Error("Player by ID does not match expected values")
	}

	foundPlayer = GetPlayerByID("3")
	if foundPlayer != nil {
		t.Error("Player by ID (Missing) does not match expected values")
	}

	foundPlayer = GetPlayerByToken("TOKEN", "UUID")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo)) {
		t.Error("Player by Token does not match expected values")
	}

	foundPlayer = GetPlayerByToken("INVALID", "UUID")
	if foundPlayer != nil {
		t.Error("Player by Token (Missing) does not match expected values")
	}

	foundPlayer = GetPlayerByToken("TOKEN", "INVALID")
	if foundPlayer != nil {
		t.Error("Player by Token (Missing) does not match expected values")
	}

	foundPlayer = GetPlayerByToken("INVALID", "INVALID")
	if foundPlayer != nil {
		t.Error("Player by Token (Missing) does not match expected values")
	}
}

func TestAddPlayers(t *testing.T) {
	ClearPlayerTable()

	var playerInfo1 = structs.PlayerInfo{
		ID: "1",
		Name: structs.PlayerName{
			Adjective: 1,
			Noun:      2,
		},
		Picture: 3,
	}

	var playerInfo2 = structs.PlayerInfo{
		ID: "2",
		Name: structs.PlayerName{
			Adjective: 99,
			Noun:      98,
		},
		Picture: 97,
	}

	//add players
	AddPlayer("TOKEN_ONE", "UUID_ONE", playerInfo1)
	AddPlayer("TOKEN_TWO", "UUID_TWO", playerInfo2)

	//initialize test variables
	var foundPlayer *structs.PlayerInfo

	//check if getters work
	foundPlayer = GetPlayerByID("1")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo1)) {
		t.Error("Player1 by ID does not match expected values")
	}

	foundPlayer = GetPlayerByID("2")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo2)) {
		t.Error("Player2 by ID does not match expected values")
	}

	foundPlayer = GetPlayerByToken("TOKEN_ONE", "UUID_ONE")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo1)) {
		t.Error("Player1 by Token does not match expected values")
	}

	foundPlayer = GetPlayerByToken("TOKEN_TWO", "UUID_TWO")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo2)) {
		t.Error("Player2 by Token does not match expected values")
	}
}

func TestAddPlayerDuplicateToken(t *testing.T) {
	ClearPlayerTable()

	var playerInfo1 = structs.PlayerInfo{
		ID: "1",
		Name: structs.PlayerName{
			Adjective: 1,
			Noun:      2,
		},
		Picture: 3,
	}

	var playerInfo2 = structs.PlayerInfo{
		ID: "2",
		Name: structs.PlayerName{
			Adjective: 99,
			Noun:      98,
		},
		Picture: 97,
	}

	//initialize test variables
	var success bool
	var foundPlayer *structs.PlayerInfo

	//add one player
	success = AddPlayer("TOKEN", "UUID_ONE", playerInfo1)
	if !success {
		t.Error("Player add failed")
	}

	//add duplicate player
	success = AddPlayer("TOKEN", "UUID_TWO", playerInfo2)
	if success {
		t.Error("Duplicate Token was added")
	}

	//check if getters work
	foundPlayer = GetPlayerByID("1")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo1)) {
		t.Error("Player by ID does not match expected values")
	}

	foundPlayer = GetPlayerByID("2")
	if foundPlayer != nil {
		t.Error("Player by ID (Missing) does not match expected values", foundPlayer)
	}

	foundPlayer = GetPlayerByToken("TOKEN", "UUID_ONE")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo1)) {
		t.Error("Player by Token does not match expected values")
	}

	foundPlayer = GetPlayerByToken("TOKEN", "UUID_TWO")
	if foundPlayer != nil {
		t.Error("Player by Token (Missing) does not match expected values")
	}
}

func TestAddPlayerDuplicateUUID(t *testing.T) {
	ClearPlayerTable()

	var playerInfo1 = structs.PlayerInfo{
		ID: "1",
		Name: structs.PlayerName{
			Adjective: 1,
			Noun:      2,
		},
		Picture: 3,
	}

	var playerInfo2 = structs.PlayerInfo{
		ID: "2",
		Name: structs.PlayerName{
			Adjective: 99,
			Noun:      98,
		},
		Picture: 97,
	}

	//initialize test variables
	var success bool
	var foundPlayer *structs.PlayerInfo

	//add one player
	success = AddPlayer("TOKEN_ONE", "UUID", playerInfo1)
	if !success {
		t.Error("Player add failed")
	}

	//add duplicate player
	success = AddPlayer("TOKEN_TWO", "UUID", playerInfo2)
	if !success {
		t.Error("Player add failed")
	}

	//check if getters work
	foundPlayer = GetPlayerByID("1")
	if foundPlayer != nil {
		t.Error("Player by ID (Missing) does not match expected values")
	}

	foundPlayer = GetPlayerByID("2")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo2)) {
		t.Error("Player by ID does not match expected values", foundPlayer)
	}

	foundPlayer = GetPlayerByToken("TOKEN_ONE", "UUID")
	if foundPlayer != nil {
		t.Error("Player by Token (Missing) does not match expected values")
	}

	foundPlayer = GetPlayerByToken("TOKEN_TWO", "UUID")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo2)) {
		t.Error("Player by Token does not match expected values")
	}
}

func TestAddPlayerDuplicateID(t *testing.T) {
	ClearPlayerTable()

	var playerInfo1 = structs.PlayerInfo{
		ID: "1",
		Name: structs.PlayerName{
			Adjective: 1,
			Noun:      2,
		},
		Picture: 3,
	}

	var playerInfo2 = structs.PlayerInfo{
		ID: "1",
		Name: structs.PlayerName{
			Adjective: 99,
			Noun:      98,
		},
		Picture: 97,
	}

	//initialize test variables
	var success bool
	var foundPlayer *structs.PlayerInfo

	//add one player
	success = AddPlayer("TOKEN_ONE", "UUID_ONE", playerInfo1)
	if !success {
		t.Error("Player add failed")
	}

	//add duplicate player
	success = AddPlayer("TOKEN_TWO", "UUID_TWO", playerInfo2)
	if success {
		t.Error("Duplicate ID was added")
	}

	//check if getters work
	foundPlayer = GetPlayerByID("1")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo1)) {
		t.Error("Player by ID does not match expected values")
	}

	foundPlayer = GetPlayerByToken("TOKEN_ONE", "UUID_ONE")
	if !(reflect.DeepEqual(*foundPlayer, playerInfo1)) {
		t.Error("Player by Token does not match expected values")
	}

	foundPlayer = GetPlayerByToken("TOKEN_TWO", "UUID_TWO")
	if foundPlayer != nil {
		t.Error("Player by Token (Missing) does not match expected values")
	}
}

func TestUpdatePlayer(t *testing.T) {
	ClearPlayerTable()

	var playerInfo = structs.PlayerInfo{
		ID: "1",
		Name: structs.PlayerName{
			Adjective: 1,
			Noun:      2,
		},
		Picture: 3,
	}

	var playerInfoNew = structs.PlayerInfo{
		ID: "1",
		Name: structs.PlayerName{
			Adjective: 99,
			Noun:      98,
		},
		Picture: 97,
	}

	//add one player
	AddPlayer("TOKEN", "UUID", playerInfo)

	//initialize test variables
	var foundPlayer *structs.PlayerInfo

	//update player
	UpdatePlayer("TOKEN", &playerInfoNew)

	//check player value is still correct
	foundPlayer = GetPlayerByID("1")
	if !(reflect.DeepEqual(*foundPlayer, playerInfoNew)) {
		if (reflect.DeepEqual(*foundPlayer, playerInfo)) {
			t.Error("Player data was not updated")
		} else {
			t.Error("Player data was changed to incorrect values")
		}
	}

}
