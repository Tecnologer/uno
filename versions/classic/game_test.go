package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/tecnologer/uno/engine"
)

func Test_classic_Shuffle(t *testing.T) {
	type fields struct {
		drawPile []engine.Card
		players  []engine.Player
	}

	cards, _ := loadCards()
	type args struct {
		times int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "shuffle_2",
			fields: fields{
				drawPile: append([]engine.Card{}, cards...),
			},
			args: args{
				times: 2,
			},
		},
		{
			name: "shuffle_random",
			fields: fields{
				drawPile: append([]engine.Card{}, cards...),
			},
			args: args{
				times: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				drawPile: tt.fields.drawPile,
				players:  tt.fields.players,
			}
			c.Shuffle(tt.args.times)

			data, err := json.MarshalIndent(c.drawPile, "", "  ")
			if err != nil {
				t.Error(err)
			}

			err = ioutil.WriteFile(fmt.Sprintf("%s_test.json", tt.name), data, 0644)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
