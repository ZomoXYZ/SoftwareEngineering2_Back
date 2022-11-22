package gameplay

import (
	"edu/letu/wan/database"
	"edu/letu/wan/structs"
	"encoding/json"
	"strconv"
)

func RunHostCommand(game *ActiveGame, cmd *PlayerCommandMessage) bool {
	switch (cmd.Cmd.Command) {
	case "kick":
		commandKick(game, cmd)
		return true

	case "setpassword":
		if !game.InLobby {
			return false
		}
		commandSetPass(game, cmd)
		return true

	case "setpointgoal":
		if !game.InLobby {
			return false
		}
		commandSetPointGoal(game, cmd)
		return true

	case "start":
		if !game.InLobby {
			return false
		}
		commandStart(game, cmd)
		return true
	}
	return false
}

func commandKick(game *ActiveGame, cmd *PlayerCommandMessage) {
	if len(cmd.Cmd.Args) < 1 {
		cmd.Player.Send <- Command("error", "kick", "no player id provided")
	}
	playerID := cmd.Cmd.Args[0]
	for _, player := range game.Players {
		if player.Player.ID == playerID {
			player.Send <- Command("kicked")
			player.Close <- true
			return
		}
	}
}

func commandSetPass(game *ActiveGame, cmd *PlayerCommandMessage) {
	if len(cmd.Cmd.Args) < 1 {
		database.UpdateLobbyPassword(game.LobbyID, "")
	}
	database.UpdateLobbyPassword(game.LobbyID, cmd.Cmd.Args[0])
}

func commandSetPointGoal(game *ActiveGame, cmd *PlayerCommandMessage) {
	if len(cmd.Cmd.Args) < 1 {
		cmd.Player.Send <- Command("error", "setpointgoal", "no point goal provided")
	}
	pointGoal, err := strconv.Atoi(cmd.Cmd.Args[0])
	if err != nil {
		cmd.Player.Send <- Command("error", "setpointgoal", "point goal is not an integer")
	}

	game.Settings.PointsToWin = pointGoal
}

type GameStartJson struct {
	Cards []structs.Card `json:"cards" binding:"required"`
	DiscardPile structs.Card `json:"discardPile" binding:"required"`
}

func commandStart(game *ActiveGame, cmd *PlayerCommandMessage) {
	game.InLobby = false

	// set game state
	game.ResetState(false)

	// fill each player's hands
	for _, player := range game.GetPlayers() {
		player.Cards = make([]structs.Card, 0)
		for i := 0; i < 5; i++ {
			player.Cards = append(player.Cards, structs.RandomCard())
		}
		cardsJson, err := json.Marshal(GameStartJson{
			Cards: player.Cards,
			DiscardPile: game.GameState.DiscardPile,
		})
		if err != nil {
			return
		}
		player.Send <- Command("starting", string(cardsJson))
	}
}