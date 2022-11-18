package gameplay

import (
	"edu/letu/wan/structs"
	"encoding/json"
	"strconv"
)

func RunPlayerCommand(game *ActiveGame, cmd *PlayerCommandMessage) bool {
	// there are no player commands that can run in the lobby
	if game.InLobby {
		return false
	}

	players := game.GetPlayers()

	playerIndex := -1
	for i, player := range players {
		if player.Player.ID == cmd.Player.Player.ID {
			playerIndex = i
			break
		}
	}
	if playerIndex == -1 {
		return false
	}

	if game.GameState.EveryoneIn && game.GameState.CurrentPlayer != playerIndex {
		return false
	}
	player := players[playerIndex]

	switch (cmd.Cmd.Command) {
	case "ingame":
		commandIngame(game, player)
		return true
	case "draw":
		if !game.TurnState.DidDraw {
			commandDraw(game, player, cmd.Cmd.Args)
		}
		return true
	case "discard":
		if game.TurnState.DidDraw && !game.TurnState.DidDiscard {
			commandDiscard(game, player, cmd.Cmd.Args)
		}
		return true
	case "play":
		if game.TurnState.DidDraw && game.TurnState.DidDiscard && !game.TurnState.DidPlay {
			commandPlay(game, player, cmd.Cmd.Args)
		}
		return true
	}
	return false
}

type PlayerDrewCard struct {
	From int `json:"from" binding:"required"`
	Card structs.Card `json:"card" binding:"required"`
}

// {
// 	"cards": my hand,
// 	"discardPile": discard pile,
// 	"points": each player points,
// 	"turn": whose turn it is,
// }

type SendTurnState struct {
	Cards []structs.Card `json:"cards" binding:"required"`
	DiscardPile structs.Card `json:"discardPile" binding:"required"`
	Points []int `json:"points" binding:"required"`
	Turn string `json:"turn" binding:"required"`
}

func broadcastTurnState(game *ActiveGame) {
	players := game.GetPlayers()
	allPlayerPoints := make([]int, len(players))
	for i, player := range players {
		allPlayerPoints[i] = player.Points
	}

	turnPlayerID := players[game.GameState.CurrentPlayer].Player.ID

	for _, player := range players {
		turnStateJSON, err := json.Marshal(SendTurnState{
			Cards: player.Cards,
			DiscardPile: game.GameState.DiscardPile,
			Points: allPlayerPoints,
			Turn: turnPlayerID,
		})
		if err != nil {
			return
		}
		player.Send <- Command("turn", string(turnStateJSON))
	}
}

func commandIngame(game *ActiveGame, player *GamePlayer) {
	player.InGame = true
	for _, player := range game.GetPlayers() {
		if !player.InGame {
			return
		}
	}
	game.GameState.EveryoneIn = true
	broadcastTurnState(game)
}

// > draw 0/1
// 0 = draw from deck
// 1 = draw from discard pile
func commandDraw(game *ActiveGame, player *GamePlayer, args []string) {
	drawType, err := strconv.Atoi(args[0]);
	if err != nil {
		player.Send <- Command("badcommand")
		return
	}

	var card structs.Card

	if drawType == 0 {
		card = structs.RandomCard()
	} else if drawType == 1 {
		card = game.GameState.DiscardPile
	} else {
		player.Send <- Command("badcommand")
		return
	}
	
	player.Cards = append(player.Cards, card)

	drewJSON, err := json.Marshal(PlayerDrewCard{
		From: drawType,
		Card: card,
	})
	if err != nil {
		return
	}
	game.TurnState.DidDraw = true
	game.Broadcast(Command("drew", string(drewJSON)))
}

// > discard cardType
func commandDiscard(game *ActiveGame, player *GamePlayer, args []string) {
	cardType, err := strconv.Atoi(args[0]);
	if err != nil {
		player.Send <- Command("badcommand")
		return
	}

	var discardCard = structs.Card(cardType)
	var found = false

	for i, card := range player.Cards {
		if card == discardCard {
			player.Cards = append(player.Cards[:i], player.Cards[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		player.Send <- Command("badcommand")
		return
	}

	game.GameState.DiscardPile = discardCard
	game.Broadcast(Command("discarded", strconv.Itoa(int(discardCard))))
}

type CardsBodyJSON struct {
	Cards []structs.Card `json:"cards" binding:"required"`
}


// > play {"cards": []cardType}
func commandPlay(game *ActiveGame, player *GamePlayer, args []string) {
	if len(args) == 0 {
		// no args, player is passing
		game.Broadcast(Command("passed"))
	} else {
		// args, player is playing cards
		var playData CardsBodyJSON
		err := json.Unmarshal([]byte(args[0]), &playData)
		if err != nil {
			player.Send <- Command("badcommand")
			return
		}

		// update player's cards
		playerCards := player.Cards
		for _, card := range playData.Cards {
			var found = false
			for i, playerCard := range playerCards {
				if playerCard == card {
					playerCards = append(playerCards[:i], playerCards[i+1:]...)
					found = true
					break
				}
			}

			if !found {
				player.Send <- Command("badcommand")
				return
			}
		}
		player.Cards = playerCards

		// TODO count points and confirm cards are valid

		// re-marshal playData and broadcast
		playDataJSON, err := json.Marshal(playData)
		if err != nil {
			return
		}
		game.TurnState.DidDraw = true
		game.Broadcast(Command("played", string(playDataJSON)))

		// check if player won
		if player.Points >= game.Settings.PointsToWin {
			game.Broadcast(Command("gameover", player.Player.ID))
			return
		}

		// draw cards until player has 5 cards
		drewCards := make([]structs.Card, 0)
		for len(player.Cards) < 5 {
			card := structs.RandomCard()
			player.Cards = append(player.Cards, card)
			drewCards = append(drewCards, card)
		}
		drewCardsJson, err := json.Marshal(CardsBodyJSON{
			Cards: drewCards,
		})
		if err != nil {
			return
		}
		player.Send <- Command("autodraw", string(drewCardsJson))
	}

	// update game state
	game.GameState.CurrentPlayer = (game.GameState.CurrentPlayer + 1) % len(game.GetPlayers())
	game.TurnState = TurnState{
		DidDraw: false,
		DidDiscard: false,
		DidPlay: false,
	}
	broadcastTurnState(game)
}