package gameplay

import (
	"edu/letu/wan/database"
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