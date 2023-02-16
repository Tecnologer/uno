package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/tecnologer/uno/engine"
	"github.com/tecnologer/uno/plugins"
)

var (
	options = []string{
		"List games",
		"Play Game",
	}

	availableGames []engine.Game
)

func main() {
	err := loadAvailableGames()
	if err != nil {
		log.Fatal("loading available games", err)
	}

	displayOptions()
}

func loadAvailableGames() (err error) {
	availableGames, err = plugins.Load()
	if err != nil {
		return err
	}

	return nil
}

func displayOptions() {
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	selectOptions()
}

func selectOptions() {
	var option int
	_, e := fmt.Scanf("%d", &option)
	if e != nil {
		fmt.Printf("invalid option (%v). try again.", e)
		selectOptions()
		return
	}

	if option < 1 || option >= len(options) {
		fmt.Printf("invalid option. try again.")
		selectOptions()
		return
	}

	switch options[option-1] {
	case "List games":
		displayAvailableGames()
	case "Play Game":
		playGame()
	}
}

func displayAvailableGames() {
	for i, game := range availableGames {
		fmt.Printf("\t+ %d) %s\n", i+1, game.GetMetadata())
	}
}

func playGame() {
	displayAvailableGames()

	var option string
	_, e := fmt.Scanf("%s", &option)
	if e != nil {
		fmt.Printf("invalid game input (%v). try again.", e)
		playGame()
		return
	}

	var game engine.Game
	for _, g := range availableGames {
		if strings.EqualFold(g.GetMetadata().GetName(), option) {
			game = g
			break
		}
	}

	if game == nil {
		fmt.Printf("invalid game: %s. try again.", option)
		playGame()
		return
	}

	// uno.StartGame(game)
}
