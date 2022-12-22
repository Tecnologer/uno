package game

import (
	"github.com/tecnologer/uno/src/card"
	"github.com/tecnologer/uno/src/player"
)

type Game interface {
	Start() (card.Card, player.Player)
	Shuffle(int)
	GetDeck() []card.Card
	GetPlayers() []player.Player
	GetCurrentPlayer() player.Player
	SetNextPlayer(player.Player)
	PlayCard(player.Player, card.Card)
	GetDirection() string
	SetDirection(string)
	SayUno(player.Player)
	GetCardFromDeck() card.Card
}
