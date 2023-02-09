package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/tecnologer/uno/src/engine"
)

const (
	reverseValue   = "reverse"
	skipValue      = "skip"
	draw2Value     = "+2"
	draw4Value     = "+4"
	undefinedColor = "undefinded_color"
	wildValue      = "wild"
	draw4WildValue = "+4,wild"

	blackColor  = "black"
	yellowColor = "yellow"
	redColor    = "red"
	blueColor   = "blue"
	greenColor  = "green"
)

var (
	numberReg = regexp.MustCompile(`^\d$`)
	drawReg   = regexp.MustCompile(`^\+\d(,wild)?$`)
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
	if c.isWild() {
		return blackColor
	}

	return c.Color
}

func (c *card) Play(currentColor string, d []engine.Card) (engine.Result, error) {
	if canPlay, err := c.CanPlay(currentColor, d); !canPlay {
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

func (c *card) CanPlay(currentColor string, d []engine.Card) (bool, error) {
	if c.isWild() {
		return true, nil
	}

	return c.canPlayColor(currentColor, d)
}

func (c *card) canPlayColor(currentColor string, d []engine.Card) (canPlay bool, _ error) {
	lastCard := d[len(d)-1].(*card)

	if lastCard.isWild() {
		canPlay = c.GetColor() == currentColor
		if !canPlay {
			return false, fmt.Errorf("invalid color (expected: %s, got: %s)",
				currentColor,
				c.GetColor(),
			)
		}

		return true, nil
	}

	canPlay = lastCard.GetColor() == c.GetColor() ||
		lastCard.GetValue() == c.GetValue()

	if !canPlay {
		return false, fmt.Errorf("invalid color (expected: %s, got: %s) or value (expected: %s, got: %s)",
			lastCard.GetColor(),
			c.GetColor(),
			lastCard.GetValue(),
			lastCard.GetValue(),
		)
	}

	return true, nil
}

func (c *card) isWild() bool {
	return strings.Contains(c.GetValue(), "wild")
}

func (c *card) isNumber() bool {
	return numberReg.MatchString(c.GetValue())
}

func (c *card) isReverse() bool {
	return strings.EqualFold(c.GetValue(), reverseValue)
}

func (c *card) isSkip() bool {
	return strings.EqualFold(c.GetValue(), skipValue) || c.getDrawCount() > 0
}

func (c *card) getNextColor() string {
	if c.isWild() {
		return undefinedColor
	}

	return c.GetColor()
}

func (c *card) getDrawCount() int {
	value := c.GetValue()
	if strings.EqualFold(value, draw2Value) {
		return 2
	}

	if strings.Contains(value, draw4Value) {
		return 4
	}

	return 0
}

func (c *card) isDraw() bool {
	return drawReg.MatchString(c.GetValue())
}

//isValid returns true if the card is valid, usefull for load from json
func (c *card) isValid() bool {
	return (c.isNumber() ||
		c.isReverse() ||
		c.isSkip() ||
		c.isWild() ||
		c.isDraw()) && c.isColorValid()
}

//isColorValid checks if the card has the correct color based on its value
func (c *card) isColorValid() bool {
	switch {
	case c.isNumber() || //is number between 0 and 9
		c.isReverse() || //is reverse card
		(c.isSkip() && !c.isWild()) || //is skip card (it ignored if is skip by draw card)
		(c.isDraw() && !c.isWild()): //is draw card (it ignored if is wild card)
		return c.GetColor() == blueColor ||
			c.GetColor() == redColor ||
			c.GetColor() == yellowColor ||
			c.GetColor() == greenColor
	case c.isWild():
		return c.GetColor() == blackColor
	default:
		return false
	}
}
