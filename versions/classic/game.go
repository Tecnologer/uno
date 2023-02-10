package main

import (
	"fmt"

	"github.com/tecnologer/uno/engine"
	"github.com/tecnologer/uno/src/tools/cards"
)

const (
	left2Rigth = "left_rigth"
	rigth2Left = "rigth_left"
)

type classic struct {
	metadata      metadata
	direction     string
	currentColor  string
	currentPlayer engine.Player
	playerSaidUno engine.Player
	drawPile      []engine.Card
	discardedPile []engine.Card
	players       []engine.Player
	output        chan engine.Result
}

func New() classic {
	return classic{}
}

func (c *classic) GetMetadata() engine.Metadata {
	return c.metadata
}

func (c *classic) New() (_ chan engine.Result, err error) {
	if c.output != nil {
		close(c.output)
	}

	c.drawPile, err = loadCards()
	if err != nil {
		return nil, err
	}

	c.direction = rigth2Left

	c.output = make(chan engine.Result)

	return c.output, nil
}

func (c *classic) Start() (engine.Card, engine.Player, error) {
	if len(c.players) < 2 {
		return nil, nil, fmt.Errorf("it required at least 2 players to start")
	}

	player := c.players[0]
	card := c.DrawCard()

	return card, player, nil
}

func (c *classic) DrawCard() engine.Card {
	card := c.drawPile[len(c.drawPile)-1]
	c.drawPile = c.drawPile[:len(c.drawPile)-1]

	return card
}

func (c *classic) Shuffle(times int) {
	c.drawPile = cards.Shuffle(c.drawPile, times)
}

func (g *classic) GetDrawPile() []engine.Card {
	return g.drawPile
}

func (c *classic) GetPlayers() []engine.Player {
	return c.players
}

func (c *classic) GetCurrentPlayer() engine.Player {
	return c.currentPlayer
}

func (c *classic) SetNextPlayer(nextPlayer engine.Player) {
	c.currentPlayer = nextPlayer
}

func (c *classic) PlayCard(player engine.Player, card engine.Card) error {
	result, err := card.Play(c.currentColor, c.discardedPile)
	if err != nil {
		return err
	}

	c.output <- result

	return nil
}

func (c *classic) GetDirection() string {
	return c.direction
}

func (c *classic) SetDirection(direction string) {
	c.direction = direction
}

func (c *classic) SayUno(player engine.Player) {
	c.playerSaidUno = player
}
func (c *classic) GetDiscardedPile() []engine.Card {
	return c.discardedPile
}
