package player

import "github.com/tecnologer/uno/src/card"

type Player interface {
	GetCards() []card.Card
	GetIndex() int
}
