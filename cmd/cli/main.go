package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tecnologer/uno"
	"github.com/tecnologer/uno/engine"
	"github.com/tecnologer/uno/plugins"
)

var (
	options        []option
	availableGames []engine.Game

	currentGame engine.Game
)

// option struct
type option struct {
	name        string
	f           func()
	description string
}

func init() {
	options = []option{
		{
			name:        "List games",
			f:           displayAvailableGames,
			description: "list available games",
		},
		{
			name:        "Play",
			f:           playGame,
			description: "play a game",
		},
		{
			name:        "Add player",
			f:           addPlayer,
			description: "add a player to the current game",
		},
		{
			name:        "Remove player",
			f:           removePlayer,
			description: "remove a player from the current game",
		},
		{
			name:        "Start",
			f:           startGame,
			description: "start the current game",
		},
		{
			name:        "Exit",
			f:           exit,
			description: "exit the game",
		},
	}
}

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
	for i, op := range options {
		fmt.Printf("%d. %s: %s\n", i+1, op.name, op.description)
		i++
	}

	selectOptions()
}

func selectOptions() {
	var op int
	_, e := fmt.Scanf("%d", &op)
	if e != nil {
		fmt.Printf("invalid op (%v). try again.", e)
		selectOptions()
		return
	}

	fmt.Printf("op: %d\n", op)

	if op < 1 || op > len(options) {
		fmt.Printf("invalid op. try again. Select a number between 1 and %d\n", len(options))
		selectOptions()
		return
	}

	fmt.Println("selected option:", options[op-1].name)
	options[op-1].f()
}

func displayAvailableGames() {
	fmt.Println("Available games: ")
	for i, game := range availableGames {
		fmt.Printf("\t+ %d) %s\n", i+1, game.GetMetadata())
	}
}

func playGame() {
	displayAvailableGames()

	fmt.Println("Select a game: ")
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

	err := uno.CreateNewGame(game.GetMetadata().GetName(), game)
	if err != nil {
		log.Println("creating new game", err)
		return
	}

	err = uno.StartGame(game.GetMetadata().GetName())
	if err != nil {
		log.Println("starting game", err)
		return
	}

	currentGame = game
}

// addPlayer adds a player to the game
func addPlayer() {
	//check if there is a game
	if currentGame == nil {
		fmt.Println("there isn't a game. try again.")
		displayOptions()
		return
	}

	//read name of the player
	var name string
	fmt.Print("enter the name of the player: ")
	_, e := fmt.Scanf("%s", &name)
	if e != nil {
		fmt.Printf("invalid name (%v). try again.", e)
		addPlayer()
		return
	}

	//add player to the current game
	err := uno.RegisterPlayer(currentGame.GetMetadata().GetName(), name)
	if err != nil {
		fmt.Printf("error adding player (%v). try again.", err)
		addPlayer()
		return
	}

	//display the player added correctly
	fmt.Printf("player %s added correctly", name)

	//ask if the user wants to add another player
	fmt.Print("do you want to add another player? (y/n): ")
	var option string
	_, e = fmt.Scanf("%s", &option)
	if e != nil {
		fmt.Printf("invalid option (%v). try again.", e)
		displayOptions()
		return
	}

	if strings.EqualFold(option, "y") {
		addPlayer()
		return
	}
}

func exit() {
	fmt.Println("bye!")
	os.Exit(0)
}

func removePlayer() {
	//TODO: implement remove player
}

func startGame() {
	//check if there is a game
	if currentGame == nil {
		fmt.Println("there isn't a game. try again.")
		displayOptions()
		return
	}

	//start the game
	err := uno.StartGame(currentGame.GetMetadata().GetName())
	if err != nil {
		fmt.Printf("error starting game (%v). try again.", err)
		startGame()
		return
	}

	//display the game started correctly
	fmt.Printf("game %s started correctly", currentGame.GetMetadata().GetName())
}
