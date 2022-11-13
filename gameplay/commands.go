package gameplay

func hostCommand(game *ActiveGame, cmd *PlayerCommandMessage) bool {
	//TODO
	switch (cmd.Cmd.Command) {
	case "hi":
		cmd.Player.send <- Command("hi", "host")
		return true
	}
	return false
}

func playerCommand(game *ActiveGame, cmd *PlayerCommandMessage) bool {
	//TODO
	switch (cmd.Cmd.Command) {
	case "hi":
		cmd.Player.send <- Command("hi", "player")
		return true
	}
	return false
}