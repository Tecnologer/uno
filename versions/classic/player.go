package main

import "github.com/tecnologer/uno/engine"

type player struct {
	name  string
	cards []engine.Card
	index int
}

func (p *player) GetCards() []engine.Card {
	return p.cards
}

func (p *player) GetIndex() int {
	return p.index
}

// GetName returns the player name
func (p *player) GetName() string {
	return p.name
}
