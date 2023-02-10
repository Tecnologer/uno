package engine

type Card interface {
	GetValue() string
	GetColor() string
	Play(string, []Card) (Result, error)
	CanPlay(string, []Card) (bool, error)
}
