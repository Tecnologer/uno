package uno

import (
	"log"

	"github.com/tecnologer/uno/engine"
)

const (
	leftDirection  = "left"
	rigthDirection = "rigth"
)

func StartGame(game engine.Game) {
	firstCard, firstPlayer, err := game.Start()
	if err != nil {
		log.Fatal(err)
	}

	game.SetDirection(leftDirection)
	game.Shuffle(0)
	game.PlayCard(firstPlayer, firstCard)
	game.SetNextPlayer(firstPlayer)
}
