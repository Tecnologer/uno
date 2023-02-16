package main

import (
	"fmt"

	"github.com/tecnologer/uno/engine"
	"github.com/tecnologer/uno/src/tools/cards"
)

type direction string

const (
	left2Rigth direction = "left_rigth"
	rigth2Left direction = "rigth_left"

	cardsPerPlayer int = 7
)

func (d direction) IsLeft() bool {
	return d == left2Rigth
}

type classic struct {
	metadata      metadata
	direction     direction
	currentColor  string
	currentPlayer engine.Player
	playerSaidUno engine.Player
	drawPile      []engine.Card
	discardedPile []engine.Card
	players       []engine.Player
}

func New() classic {
	return classic{}
}

func (c *classic) GetMetadata() engine.Metadata {
	return c.metadata
}

func (c *classic) New() (err error) {
	c.drawPile, err = loadCards()
	if err != nil {
		return err
	}

	c.direction = rigth2Left

	return nil
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

func (c *classic) AddPlayer(playerName string) error {
	if c.isPlayerPlaying(playerName) {
		return fmt.Errorf("the player %s is in the game", playerName)
	}

	player := &player{
		name:   playerName,
		output: make(chan engine.TurnResult, 1),
		cards:  c.getCardsForPlayer(),
		index:  len(c.players),
	}

	c.players = append(c.players, player)

	return nil
}

func (c *classic) isPlayerPlaying(playerName string) bool {
	for _, player := range c.players {
		if player.GetName() == playerName {
			return true
		}
	}

	return false
}

func (c *classic) getCardsForPlayer() []engine.Card {
	cards := make([]engine.Card, cardsPerPlayer)

	for i := range cards {
		cards[i] = c.DrawCard()
	}

	return cards
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

func (c *classic) PlayCard(player engine.Player, card engine.Card) (engine.TurnResult, error) {
	result, err := card.Play(c.currentColor, c.discardedPile)
	if err != nil {
		return nil, err
	}

	c.updateDirection(result.IsReverse())

	r := &turnResult{
		card:      card,
		player:    c.getNextPlayer(player, result.IsReverse(), result.IsSkip()),
		direction: c.direction,
	}

	return r, nil
}

func (c *classic) getNextPlayer(currentPlayer engine.Player, isReverse, isSkip bool) engine.Player {
	var (
		nextPlayer      engine.Player
		nextPlayerIndex int
	)

	if isReverse {
		nextPlayerIndex = currentPlayer.GetIndex() - 1
	} else {
		nextPlayerIndex = currentPlayer.GetIndex() + 1
	}

	if nextPlayerIndex < 0 {
		nextPlayerIndex = len(c.players) - 1
	}

	if nextPlayerIndex >= len(c.players) {
		nextPlayerIndex = 0
	}

	if isSkip {
		nextPlayerIndex++
	}

	nextPlayer = c.players[nextPlayerIndex]

	return nextPlayer
}

func (c *classic) updateDirection(isReverse bool) {
	if !isReverse {
		return
	}

	if c.direction == left2Rigth {
		c.direction = rigth2Left
		return
	}

	c.direction = left2Rigth
}

func (c *classic) GetDirection() engine.Direction {
	return c.direction
}

func (c *classic) SetDirection(d engine.Direction) {
	if d.IsLeft() {
		c.direction = left2Rigth
		return
	}

	c.direction = rigth2Left
}

func (c *classic) SayUno(player engine.Player) {
	c.playerSaidUno = player
}

func (c *classic) GetDiscardedPile() []engine.Card {
	return c.discardedPile
}

// func (c *classic) Close() {
// 	close(c.output)
// }
