package main

import (
	"github.com/tecnologer/uno/engine"
	"reflect"
	"testing"
)

func Test_player_GetCards(t *testing.T) {
	type fields struct {
		cards []engine.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   []engine.Card
	}{
		{
			name: "get_cards",
			fields: fields{
				cards: []engine.Card{
					&card{
						Color: redColor,
						Value: "1",
					},
				},
			},
			want: []engine.Card{
				&card{
					Color: redColor,
					Value: "1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &player{
				cards: tt.fields.cards,
			}
			if got := p.GetCards(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_player_GetIndex(t *testing.T) {
	type fields struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "get_index",
			fields: fields{
				index: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &player{
				index: tt.fields.index,
			}
			if got := p.GetIndex(); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_player_GetName(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get_name",
			fields: fields{
				name: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &player{
				name: tt.fields.name,
			}
			if got := p.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
