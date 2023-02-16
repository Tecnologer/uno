package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
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

func Test_classic_getNextPlayer(t *testing.T) {
	players := getPlayersForTests()

	type args struct {
		currentPlayer engine.Player
		isReverse     bool
		isSkip        bool
	}
	tests := []struct {
		name    string
		players []engine.Player
		args    args
		want    engine.Player
	}{
		{
			name:    "normal",
			players: players,
			args: args{
				currentPlayer: players[1],
				isReverse:     false,
				isSkip:        false,
			},
			want: players[2],
		},
		{
			name:    "skip",
			players: players,
			args: args{
				currentPlayer: players[0],
				isReverse:     false,
				isSkip:        true,
			},
			want: players[2],
		},
		{
			name:    "skip_last",
			players: players,
			args: args{
				currentPlayer: players[len(players)-1],
				isReverse:     false,
				isSkip:        true,
			},
			want: players[1],
		},
		{
			name:    "reverse",
			players: players,
			args: args{
				currentPlayer: players[2],
				isReverse:     true,
				isSkip:        false,
			},
			want: players[1],
		},
		{
			name:    "reverse_first",
			players: players,
			args: args{
				currentPlayer: players[0],
				isReverse:     true,
				isSkip:        false,
			},
			want: players[len(players)-1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				players: tt.players,
			}
			if got := c.getNextPlayer(tt.args.currentPlayer, tt.args.isReverse, tt.args.isSkip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("classic.getNextPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classic_updateDirection(t *testing.T) {
	type args struct {
		isReverse bool
	}
	tests := []struct {
		name      string
		direction direction
		args      args
		want      direction
	}{
		{
			name:      "left_to_rigth_normal",
			direction: left2Rigth,
			args: args{
				isReverse: false,
			},
			want: left2Rigth,
		},
		{
			name:      "left_to_rigth_reverse",
			direction: left2Rigth,
			args: args{
				isReverse: true,
			},
			want: rigth2Left,
		},
		{
			name:      "rigth_to_left_normal",
			direction: rigth2Left,
			args: args{
				isReverse: false,
			},
			want: rigth2Left,
		},
		{
			name:      "rigth_to_left_reverse",
			direction: rigth2Left,
			args: args{
				isReverse: true,
			},
			want: left2Rigth,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				direction: tt.direction,
			}
			c.updateDirection(tt.args.isReverse)

			if c.direction != tt.want {
				t.Errorf("classic.updateDirection() = %s, want = %s", c.direction, tt.want)
			}
		})
	}
}

func Test_classic_PlayCard(t *testing.T) {
	type fields struct {
		direction     direction
		currentColor  string
		currentPlayer engine.Player
		playerSaidUno engine.Player
		drawPile      []engine.Card
		discardedPile []engine.Card
		players       []engine.Player
	}
	type args struct {
		player engine.Player
		card   engine.Card
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    engine.TurnResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				metadata:      tt.fields.metadata,
				direction:     tt.fields.direction,
				currentColor:  tt.fields.currentColor,
				currentPlayer: tt.fields.currentPlayer,
				playerSaidUno: tt.fields.playerSaidUno,
				drawPile:      tt.fields.drawPile,
				discardedPile: tt.fields.discardedPile,
				players:       tt.fields.players,
			}
			got, err := c.PlayCard(tt.args.player, tt.args.card)
			if (err != nil) != tt.wantErr {
				t.Errorf("classic.PlayCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("classic.PlayCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getPlayersForTests() []engine.Player {
	players := make([]engine.Player, 5)
	for i := range players {
		players[i] = &player{
			name:  fmt.Sprintf("Player_%d", i),
			index: i,
		}
	}

	return players
}
