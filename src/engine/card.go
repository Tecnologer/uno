package engine

type Card interface {
	GetValue() string
	GetColor() string
	Play([]Card) (Result, error)
	CanPlay([]Card) (bool, error)
}
