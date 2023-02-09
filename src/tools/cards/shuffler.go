package cards

import (
	"math/rand"
	"time"

	"github.com/tecnologer/uno/src/engine"
)

func Shuffle(pile []engine.Card, times int) []engine.Card {
	if times <= 0 {
		return pile
	}

	if times < 0 {
		times = rand.Intn(6) + 1
	}

	positions := make(map[int]int)
	shuffledDeck := make([]engine.Card, len(pile))
	var i int

	for _, card := range pile {
		i = getRandomPos(positions, len(pile))
		positions[i]++
		shuffledDeck[i] = card
	}

	pile = shuffledDeck

	times--

	return Shuffle(pile, times)
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
