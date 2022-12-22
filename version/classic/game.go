package classic

import (
	"math/rand"
	"time"

	icard "github.com/tecnologer/uno/src/card"
	iplayer "github.com/tecnologer/uno/src/player"
)

type classic struct {
	deck    []*card
	players []*player
}

func (c *classic) Start() (icard.Card, iplayer.Player) {
	if len(c.players) == 0 {
		return nil, nil
	}

	player := c.players[0]
	card := c.GetCardFromDeck()

	return card, player
}

func (c *classic) GetCardFromDeck() icard.Card {
	card := c.deck[len(c.deck)-1]
	c.deck = c.deck[:len(c.deck)-1]

	return card
}

func (c *classic) Shuffle(times int) {
	if times == 0 {
		return
	}

	if times < 0 {
		times = rand.Intn(6) + 1
	}

	positions := make(map[int]int)
	shuffledDeck := make([]*card, len(c.deck))
	var i int

	for _, card := range c.deck {
		i = getRandomPos(positions, len(c.deck))
		positions[i]++
		shuffledDeck[i] = card
	}

	c.deck = shuffledDeck

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

// GetDeck() []card.Card
// GetPlayers() []player.Player
// GetCurrentPlayer() player.Player
// SetNextPlayer(player.Player)
// PlayCard(player.Player, card.Card)
// GetDirection() string
// 	SetDirection(string)
// 	SayUno(player.Player)
