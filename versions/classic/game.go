package classic

import (
	"math/rand"
	"time"

	"github.com/tecnologer/uno/src/engine"
)

const (
	left2Rigth = "left_rigth"
	rigth2Left = "rigth_left"
)

type classic struct {
	direction     string
	currentColor  string
	currentPlayer engine.Player
	playerSaidUno engine.Player
	drawPile      []engine.Card
	discardedPile []engine.Card
	players       []engine.Player
}

func (c *classic) Start() (engine.Card, engine.Player) {
	if len(c.players) == 0 {
		return nil, nil
	}

	player := c.players[0]
	card := c.DrawCard()

	return card, player
}

func (c *classic) DrawCard() engine.Card {
	card := c.drawPile[len(c.drawPile)-1]
	c.drawPile = c.drawPile[:len(c.drawPile)-1]

	return card
}

func (c *classic) Shuffle(times int) {
	if times <= 0 {
		return
	}

	if times < 0 {
		times = rand.Intn(6) + 1
	}

	positions := make(map[int]int)
	shuffledDeck := make([]engine.Card, len(c.drawPile))
	var i int

	for _, card := range c.drawPile {
		i = getRandomPos(positions, len(c.drawPile))
		positions[i]++
		shuffledDeck[i] = card
	}

	c.drawPile = shuffledDeck

	times--
	c.Shuffle(times)
}

func getRandomPos(taken map[int]int, limit int) int {
	n := rand.Intn(limit)
	if _, ok := taken[n]; !ok {
		return n
	}

	seed := time.Now().Unix()
	for {
		rand.Seed(seed)
		n = rand.Intn(limit)
		seed++

		if _, ok := taken[n]; !ok {
			return n
		}
	}
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

	if result.IsReverse() {

	}

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
