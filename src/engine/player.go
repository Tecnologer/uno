package engine

type Player interface {
	GetName() string
	GetCards() []Card
	GetIndex() int
}
