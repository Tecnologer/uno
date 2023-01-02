package engine

type Result interface {
	IsReverse() bool
	IsSkip() bool
	GetDrawCount() int
	GetNextColor() string
	GetDeck() []Card
}
