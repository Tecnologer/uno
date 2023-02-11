package uno

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tecnologer/uno/engine"
)

const (
	leftDirection  = "left"
	rigthDirection = "rigth"
)

var games = map[string]engine.Game{}

func CreateNewGame(gameName string, game engine.Game) error {
	if _, ok := games[gameName]; ok {
		return fmt.Errorf("there is a game with name %s", gameName)
	}

	output, err := game.New()
	if err != nil {
		return errors.Wrap(err, "uno: create new game")
	}

	games[gameName] = game

	go register(output)

	return nil
}

func StartGame(gameName string) error {
	game, ok := games[gameName]
	if !ok {
		return fmt.Errorf("there isn't a game with name %s", gameName)
	}

	minPlayers := game.GetMetadata().GetMinPlayer()
	if len(game.GetPlayers()) < minPlayers {
		return fmt.Errorf("the game requires at least %d players", minPlayers)
	}

	firstCard, firstPlayer, err := game.Start()
	if err != nil {
		return errors.Wrap(err, "uno: start game")
	}

	game.SetDirection(leftDirection)
	game.Shuffle(0)
	game.PlayCard(firstPlayer, firstCard)
	game.SetNextPlayer(firstPlayer)

	return nil
}

func register(output chan engine.Result) {
	for result := range output {

	}
}

func RegisterPlayer(gameName, playerName string) error {

}

func CloseGame(name string)
