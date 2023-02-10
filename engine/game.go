package engine

type Game interface {
	New() (chan Result, error)
	GetMetadata() Metadata
	Start() (Card, Player, error)
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
