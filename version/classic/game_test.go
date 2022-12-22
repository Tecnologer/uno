package classic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_classic_Shuffle(t *testing.T) {
	type fields struct {
		deck    []*card
		players []*player
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
				deck: append([]*card{}, cards...),
			},
			args: args{
				times: 2,
			},
		},
		{
			name: "shuffle_random",
			fields: fields{
				deck: append([]*card{}, cards...),
			},
			args: args{
				times: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &classic{
				deck:    tt.fields.deck,
				players: tt.fields.players,
			}
			c.Shuffle(tt.args.times)

			data, err := json.MarshalIndent(c.deck, "", "  ")
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
