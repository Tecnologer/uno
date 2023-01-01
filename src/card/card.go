package card

type Card interface {
	GetValue() string
	GetColor() string
	Play([]Card) ([]Card, error)
	CanPlay([]Card) (bool, error)
}
