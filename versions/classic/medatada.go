package main

import (
	"fmt"

	"github.com/tecnologer/uno/engine"
)

type metadata byte

var (
	minversion = "0"
	version    = "0"
)

func (metadata) GetName() string {
	return "Classic"
}

func (metadata) GetMaxPlayer() int {
	return 6
}

func (metadata) GetMinPlayer() int {
	return 2
}

func (metadata) GetVersion() string {
	return fmt.Sprintf("%s.%s", version, minversion)
}

func (m metadata) String() string {
	return engine.MetadataToString(m)
}
