package main

import "github.com/tecnologer/uno/engine"

type player struct {
	cards []engine.Card
	index int
}

func (p *player) GetCards() []engine.Card {
	return p.cards
}

func (p *player) GetIndex() int {
	return p.index
}
