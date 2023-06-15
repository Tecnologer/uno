package main

import (
	"github.com/tecnologer/uno/engine"
	"reflect"
	"testing"
)

func Test_result_GetDeck(t *testing.T) {
	type fields struct {
		deck []engine.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   []engine.Card
	}{
		{
			name: "get_deck",
			fields: fields{
				deck: []engine.Card{
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
			r := &result{
				deck: tt.fields.deck,
			}
			if got := r.GetDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_result_GetDrawCount(t *testing.T) {
	type fields struct {
		drawCount int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "get_draw_count",
			fields: fields{
				drawCount: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &result{
				drawCount: tt.fields.drawCount,
			}
			if got := r.GetDrawCount(); got != tt.want {
				t.Errorf("GetDrawCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_result_GetNextColor(t *testing.T) {
	type fields struct {
		nextColor string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get_next_color",
			fields: fields{
				nextColor: redColor,
			},
			want: redColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &result{
				nextColor: tt.fields.nextColor,
			}
			if got := r.GetNextColor(); got != tt.want {
				t.Errorf("GetNextColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_result_IsReverse(t *testing.T) {
	type fields struct {
		isReverse bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "is_reverse",
			fields: fields{
				isReverse: true,
			},
			want: true,
		},
		{
			name: "is_reverse_false",
			fields: fields{
				isReverse: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &result{
				isReverse: tt.fields.isReverse,
			}
			if got := r.IsReverse(); got != tt.want {
				t.Errorf("IsReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_result_IsSkip(t *testing.T) {
	type fields struct {
		isReverse bool
		isSkip    bool
		drawCount int
		nextColor string
		deck      []engine.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "is_skip",
			fields: fields{
				isSkip: true,
			},
			want: true,
		},
		{
			name: "is_skip_false",
			fields: fields{
				isSkip: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &result{
				isReverse: tt.fields.isReverse,
				isSkip:    tt.fields.isSkip,
				drawCount: tt.fields.drawCount,
				nextColor: tt.fields.nextColor,
				deck:      tt.fields.deck,
			}
			if got := r.IsSkip(); got != tt.want {
				t.Errorf("IsSkip() = %v, want %v", got, tt.want)
			}
		})
	}
}
