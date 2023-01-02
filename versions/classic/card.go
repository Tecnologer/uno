package classic

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tecnologer/uno/src/engine"
)

const (
	ReverseValue   = "reverse"
	SkipValue      = "skip"
	Draw2Value     = "+2"
	Draw4Value     = "+4"
	UndefinedColor = "undefinded_color"
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

func (c *card) Play(d []engine.Card) (engine.Result, error) {
	if canPlay, err := c.CanPlay(d); !canPlay {
		return nil, err
	}

	d = append(d, c)

	res := &result{
		deck:      d,
		isReverse: c.isReverse(),
		isSkip:    c.isSkip(),
		drawCount: c.getDrawCount(),
		nextColor: c.getNextColor(),
	}

	return res, nil
}

func (c *card) CanPlay(d []engine.Card) (bool, error) {
	if c.isWild() {
		return true, nil
	}

	return c.canPlayColor(d)
}

func (c *card) isWild() bool {
	return strings.Contains(c.GetValue(), "wild")
}

func (c *card) canPlayColor(d []engine.Card) (bool, error) {
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

func (c *card) isReverse() bool {
	return strings.EqualFold(c.GetValue(), ReverseValue)
}

func (c *card) isSkip() bool {
	return strings.EqualFold(c.GetValue(), SkipValue)
}

func (c *card) getNextColor() string {
	if c.isWild() {
		return UndefinedColor
	}

	return c.GetColor()
}

func (c *card) getDrawCount() int {
	value := c.GetValue()
	if strings.EqualFold(value, Draw2Value) {
		return 2
	}

	if strings.Contains(value, Draw4Value) {
		return 4
	}

	return 0
}
