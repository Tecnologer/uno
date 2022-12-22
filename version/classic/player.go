package classic

import (
	icard "github.com/tecnologer/uno/src/card"
)

type player struct {
	cards []icard.Card
	index int
}

func (p *player) GetCards() []icard.Card {
	return p.cards
}

func (p *player) GetIndex() int {
	return p.index
}
