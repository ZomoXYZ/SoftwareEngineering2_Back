package gameplay

func RunPlayerCommand(game *ActiveGame, cmd *PlayerCommandMessage) bool {
	//TODO
	switch (cmd.Cmd.Command) {
	case "hi":
		cmd.Player.Send <- Command("hi", "player")
		return true
	}
	return false
}