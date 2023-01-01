package classic

import (
	"fmt"
	"strconv"
	"strings"

	icard "github.com/tecnologer/uno/src/card"
)

//rawCard is a struct to load cards from json configuration
type rawCard struct {
	Value  string   `json:"value"`
	Count  int      `json:"count"`
	Colors []string `json:"colors"`
}

type card struct {
	Value string `json:"value"`
	Color string `json:"count"`
}

func (c *card) GetValue() string {
	return c.Value
}

func (c *card) GetColor() string {
	return c.Color
}

func (c *card) Play(d []icard.Card) ([]icard.Card, error) {
	if canPlay, err := c.CanPlay(d); !canPlay {
		return d, err
	}

	d = append(d, c)

	return d, nil
}

func (c *card) CanPlay(d []icard.Card) (bool, error) {
	if c.isWild() {
		return true, nil
	}

	return c.canPlayColor(d)
}

func (c *card) isWild() bool {
	return strings.Contains(c.GetValue(), "wild")
}

func (c *card) canPlayColor(d []icard.Card) (bool, error) {
	lastCard := d[len(d)-1]

	canPlay := lastCard.GetColor() == c.GetColor() ||
		lastCard.GetValue() == c.GetColor()

	if !canPlay {
		return false, fmt.Errorf("invalid color and value")
	}

	return true, nil
}

func (c *card) isNumber() bool {
	_, err := strconv.ParseInt(c.GetValue(), 10, 32)

	return err == nil
}
