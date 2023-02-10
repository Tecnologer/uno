package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/tecnologer/uno/engine"
)

func loadCards() ([]engine.Card, error) {
	dir, e := getCurrentDir()
	if e != nil {
		return nil, fmt.Errorf("classic.load: %v. get current dir", e)
	}
	cardsJsonPath := filepath.Join(dir, "cards.json")

	data, e := ioutil.ReadFile(cardsJsonPath)
	if e != nil {
		return nil, fmt.Errorf("classic.load: %v. read json file %s", e, cardsJsonPath)
	}

	var rawCards []*rawCard
	e = json.Unmarshal(data, &rawCards)
	if e != nil {
		return nil, fmt.Errorf("classic.load: %v. unmarshal json", e)
	}

	cards := make([]engine.Card, 0)
	for _, raw := range rawCards {
		for i := 0; i < raw.Count; i++ {
			for _, color := range raw.Colors {
				if raw.Value == "0-9" || raw.Value == "1-9" {
					values := strings.Split(raw.Value, "-")
					start, _ := strconv.Atoi(values[0])
					end, _ := strconv.Atoi(values[1])

					for ; start <= end; start++ {
						cards = append(cards, &card{
							Value: fmt.Sprint(start),
							Color: color,
						})
					}
					continue
				}

				cards = append(cards, &card{
					Value: raw.Value,
					Color: color,
				})
			}
		}
	}

	return cards, nil
}

//GetCallerDir returns the relative path of the file from where the function is called
func getCurrentDir() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("get dir for connection file")
	}

	return path.Dir(filename), nil
}
