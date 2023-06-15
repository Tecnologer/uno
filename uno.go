package uno

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tecnologer/uno/engine"
)

const (
	leftDirection  = "left"
	rightDirection = "right"
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

	err = game.PlayCard(firstPlayer, firstCard)
	if err != nil {
		return err
	}

	game.SetNextPlayer(firstPlayer)

	return nil
}

func register(output chan engine.Result) {
	// TODO: Implement this function
}

// RegisterPlayer registers a player in a game
func RegisterPlayer(gameName, playerName string) error {
	game, ok := games[gameName]
	if !ok {
		return fmt.Errorf("there isn't a game with name %s", gameName)
	}

	player := game.NewPlayer(playerName)
	err := game.RegisterPlayer(player)
	if err != nil {
		return errors.Wrap(err, "uno: register player")
	}

	return nil
}

// CloseGame closes the game
func CloseGame(name string) error {
	game, ok := games[name]
	if !ok {
		return fmt.Errorf("there isn't a game with name %s", name)
	}

	game.Close()

	return nil
}
