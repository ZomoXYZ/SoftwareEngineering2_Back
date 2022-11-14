package gameplay

func RunHostCommand(game *ActiveGame, cmd *PlayerCommandMessage) bool {
	//TODO
	switch (cmd.Cmd.Command) {
	case "hi":
		cmd.Player.Send <- Command("hi", "host")
		return true
	}
	return false
}