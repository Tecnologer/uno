package main

import "github.com/tecnologer/uno/engine"

type turnResult struct {
	player    engine.Player
	card      engine.Card
	direction engine.Direction
	drawCount int
}

func (r turnResult) GetCard() engine.Card {
	return r.card
}

func (r turnResult) GetPlayer() engine.Player {
	return r.player
}

func (r turnResult) GetDirection() engine.Direction {
	return r.direction
}

func (r turnResult) GetDrawCount() int {
	return r.drawCount
}
