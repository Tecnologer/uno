package main

import "github.com/tecnologer/uno/src/engine"

type result struct {
	isReverse bool
	isSkip    bool
	drawCount int
	nextColor string
	deck      []engine.Card
}

func (r *result) IsReverse() bool {
	return r.isReverse
}

func (r *result) IsSkip() bool {
	return r.isSkip
}

func (r *result) GetDrawCount() int {
	return r.drawCount
}

func (r *result) GetNextColor() string {
	return r.nextColor
}

func (r *result) GetDeck() []engine.Card {
	return r.deck
}
