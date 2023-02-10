package engine

type Metadata interface {
	GetName() string
	GetMaxPlayer() int
	GetMinPlayer() int
}
