package game

import "github.com/tecnologer/uno/src/engine"

const (
	leftDirection  = "left"
	rigthDirection = "rigth"
)

func StartGame(game engine.Game) {
	firstCard, firstPlayer := game.Start()

	game.Shuffle(0)
	game.PlayCard(nil, firstCard)
	game.SetNextPlayer(firstPlayer)
	game.SetDirection(leftDirection)
}
