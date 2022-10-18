package main

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"
	"testing"
)

func comparePlayer(player structs.PlayerInfo, id string, nameAdjective int, nameNoun int, picture int) bool {
	if player.ID != id {
		return false
	}
	if player.Name.Adjective != nameAdjective {
		return false
	}
	if player.Name.Noun != nameNoun {
		return false
	}
	if player.Picture != picture {
		return false
	}
	return true
}

// tests basic Set and Get functionality
func TestSQLTable(t *testing.T) {
	//initialize table
	database.Initialize()
	database.SetPlayer("1", "1111111111", 12, 55, 42)
	database.SetPlayer("2", "2222222222", 98, 0, 67)
	database.SetPlayer("3", "3333333333", 48, 72, 2)

	// get player 1 via ID
	var player1 = database.GetPlayer("1")

	if !(comparePlayer(*player1, "1", 12, 55, 42)) {
		t.Error("Player 1 does not match expected values")
	}

	var player2 = database.GetPlayerByToken("2222222222")

	if !(comparePlayer(*player2, "2", 98, 0, 67)) {
		t.Error("Player 2 does not match expected values")
	}

	var player3 = database.GetPlayer("3")

	if !(comparePlayer(*player3, "3", 48, 72, 2)) {
		t.Error("Player 3 does not match expected values")
	}
}