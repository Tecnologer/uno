package main

import (
	"fmt"

	"github.com/tecnologer/uno/engine"
)

type player struct {
	name   string
	output chan engine.TurnResult
	cards  []engine.Card
	index  int
}

func (p player) GetName() string {
	return p.name
}

func (p *player) GetCards() []engine.Card {
	return p.cards
}

func (p *player) GetIndex() int {
	return p.index
}

func (p *player) GetOutput() chan engine.TurnResult {
	return p.output
}

func (p *player) SendResult(r engine.TurnResult) error {
	if p.output == nil {
		return fmt.Errorf("the output channel for player %s isn't open", p.name)
	}

	p.output <- r

	return nil
}
