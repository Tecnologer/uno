package engine

import "github.com/tecnologer/uno/src/game"

const (
	leftDirection  = "left"
	rigthDirection = "rigth"
)

func StartGame(game game.Game) {
	firstCard, firstPlayer := game.Start()

	game.Shuffle(0)
	game.PlayCard(nil, firstCard)
	game.SetNextPlayer(firstPlayer)
	game.SetDirection(leftDirection)
}
