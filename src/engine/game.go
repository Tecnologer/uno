package engine

type Game interface {
	New() chan interface{}
	Start() (Card, Player)
	Shuffle(int)
	GetDiscardedPile() []Card
	GetDrawPile() []Card
	DrawCard() Card
	PlayCard(Player, Card) error
	GetPlayers() []Player
	GetCurrentPlayer() Player
	SetNextPlayer(Player)
	SayUno(Player)
	GetDirection() string
	SetDirection(string)
}
