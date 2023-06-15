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

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want classic
	}{
		{
			name: "new",
			want: classic{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classic_Close(t *testing.T) {
	type fields struct {
		output chan engine.Result
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "close",
			fields: fields{
				output: make(chan engine.Result),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				output: tt.fields.output,
			}
			c.Close()
		})
	}
}

func Test_classic_DrawCard(t *testing.T) {
	type fields struct {
		drawPile []engine.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   engine.Card
	}{
		{
			name: "draw_card",
			fields: fields{
				drawPile: []engine.Card{
					&card{
						Color: "red",
						Value: "1",
					},
				},
			},
			want: &card{
				Color: "red",
				Value: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				drawPile: tt.fields.drawPile,
			}
			if got := c.DrawCard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DrawCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classic_GetCurrentPlayer(t *testing.T) {
	type fields struct {
		currentPlayer engine.Player
	}
	tests := []struct {
		name   string
		fields fields
		want   engine.Player
	}{
		{
			name: "get_current_player",
			fields: fields{
				currentPlayer: &player{
					name: "test",
				},
			},
			want: &player{
				name: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				currentPlayer: tt.fields.currentPlayer,
			}
			if got := c.GetCurrentPlayer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classic_GetDirection(t *testing.T) {
	type fields struct {
		direction string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get_direction",
			fields: fields{
				direction: "clockwise",
			},
			want: "clockwise",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				direction: tt.fields.direction,
			}
			if got := c.GetDirection(); got != tt.want {
				t.Errorf("GetDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classic_GetDiscardedPile(t *testing.T) {
	type fields struct {
		discardedPile []engine.Card
	}

	tests := []struct {
		name   string
		fields fields
		want   []engine.Card
	}{
		{
			name: "get_discarded_pile",
			fields: fields{
				discardedPile: []engine.Card{
					&card{
						Color: "red",
						Value: "1",
					},
				},
			},
			want: []engine.Card{
				&card{
					Color: "red",
					Value: "1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				discardedPile: tt.fields.discardedPile,
			}
			if got := c.GetDiscardedPile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDiscardedPile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classic_GetDrawPile(t *testing.T) {
	type fields struct {
		drawPile []engine.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   []engine.Card
	}{
		{
			name: "get_draw_pile",
			fields: fields{
				drawPile: []engine.Card{
					&card{
						Color: "red",
						Value: "1",
					},
				},
			},
			want: []engine.Card{
				&card{
					Color: "red",
					Value: "1",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				drawPile: tt.fields.drawPile,
			}
			if got := c.GetDrawPile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDrawPile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classic_GetMetadata(t *testing.T) {
	type fields struct {
		metadata metadata
	}
	tests := []struct {
		name   string
		fields fields
		want   engine.Metadata
	}{
		{
			name: "get_metadata",
			fields: fields{
				metadata: 0,
			},
			want: metadata(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				metadata: tt.fields.metadata,
			}
			if got := c.GetMetadata(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classic_GetPlayers(t *testing.T) {
	type fields struct {
		players []engine.Player
	}
	tests := []struct {
		name   string
		fields fields
		want   []engine.Player
	}{
		{
			name: "get_players",
			fields: fields{
				players: []engine.Player{
					&player{
						name: "player1",
					},
				},
			},
			want: []engine.Player{
				&player{
					name: "player1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				players: tt.fields.players,
			}
			if got := c.GetPlayers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlayers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classic_NewPlayer(t *testing.T) {
	tests := []struct {
		name       string
		playerName string
		want       engine.Player
	}{
		{
			name:       "new_player",
			playerName: "test",
			want:       &player{name: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{}

			if got := c.NewPlayer(tt.playerName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_classic_SayUno(t *testing.T) {
	type fields struct {
		playerSaidUno engine.Player
	}
	type args struct {
		player engine.Player
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "say_uno",
			fields: fields{
				playerSaidUno: &player{
					name: "player1",
				},
			},
			args: args{
				player: &player{
					name: "player1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				playerSaidUno: tt.fields.playerSaidUno,
			}
			c.SayUno(tt.args.player)
		})
	}
}

func Test_classic_SetDirection(t *testing.T) {
	type fields struct {
		direction string
	}
	type args struct {
		direction string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "set_direction",
			fields: fields{
				direction: "clockwise",
			},
			args: args{
				direction: "counterclockwise",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				direction: tt.fields.direction,
			}
			c.SetDirection(tt.args.direction)
		})
	}
}

func Test_classic_SetNextPlayer(t *testing.T) {
	type fields struct {
		currentPlayer engine.Player
		players       []engine.Player
	}
	type args struct {
		nextPlayer engine.Player
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "set_next_player",
			fields: fields{
				currentPlayer: &player{
					name: "player1",
				},
				players: []engine.Player{
					&player{
						name: "player1",
					},
					&player{
						name: "player2",
					},
				},
			},
			args: args{
				nextPlayer: &player{
					name: "player2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				currentPlayer: tt.fields.currentPlayer,
				players:       tt.fields.players,
			}
			c.SetNextPlayer(tt.args.nextPlayer)
		})
	}
}

func Test_classic_Start(t *testing.T) {
	type fields struct {
		drawPile []engine.Card
		players  []engine.Player
	}
	tests := []struct {
		name    string
		fields  fields
		want    engine.Card
		want1   engine.Player
		wantErr bool
	}{
		{
			name: "start_one_player",
			fields: fields{
				drawPile: []engine.Card{
					&card{
						Color: "red",
						Value: "1",
					},
				},
				players: []engine.Player{
					&player{
						name: "player1",
					},
				},
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "start_two_player",
			fields: fields{
				drawPile: []engine.Card{
					&card{
						Color: "red",
						Value: "1",
					},
				},
				players: []engine.Player{
					&player{
						name: "player1",
					},
					&player{
						name: "player2",
					},
				},
			},
			want: &card{
				Color: "red",
				Value: "1",
			},
			want1: &player{
				name: "player1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				drawPile: tt.fields.drawPile,
				players:  tt.fields.players,
			}
			got, got1, err := c.Start()
			if (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Start() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Start() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
