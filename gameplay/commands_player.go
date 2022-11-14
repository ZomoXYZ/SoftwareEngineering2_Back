package gameplay

func RunPlayerCommand(game *ActiveGame, cmd *PlayerCommandMessage) bool {
	// there are no player commands that can run in the lobby
	if game.InLobby {
		return false
	}

	// TODO
	switch (cmd.Cmd.Command) {
	case "hi":
		cmd.Player.Send <- Command("hi", "player")
		return true
	}
	return false
}