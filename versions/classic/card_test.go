package main

import (
	"reflect"
	"testing"

	"github.com/tecnologer/uno/src/engine"
)

func Test_card_GetValue(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "zero_red",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			want: "0",
		},
		{
			name: "reverse_blue",
			fields: fields{
				Value: reverseValue,
				Color: blueColor,
			},
			want: reverseValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			if got := c.GetValue(); got != tt.want {
				t.Errorf("card.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_GetColor(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "zero_red",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			want: redColor,
		},
		{
			name: "reverse_blue",
			fields: fields{
				Value: reverseValue,
				Color: blueColor,
			},
			want: blueColor,
		},
		{
			name: "wild_card",
			fields: fields{
				Value: wildValue,
			},
			want: blackColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			if got := c.GetColor(); got != tt.want {
				t.Errorf("card.GetColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_getDrawCount(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "zero_red",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			want: 0,
		},
		{
			name: "draw_two",
			fields: fields{
				Value: draw2Value,
				Color: redColor,
			},
			want: 2,
		},
		{
			name: "wild_card",
			fields: fields{
				Value: wildValue,
			},
			want: 0,
		},
		{
			name: "draw_four_wild",
			fields: fields{
				Value: draw4WildValue,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			if got := c.getDrawCount(); got != tt.want {
				t.Errorf("card.getDrawCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_getNextColor(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "zero_red",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			want: redColor,
		},
		{
			name: "draw_two",
			fields: fields{
				Value: draw2Value,
				Color: redColor,
			},
			want: redColor,
		},
		{
			name: "wild_card",
			fields: fields{
				Value: wildValue,
			},
			want: undefinedColor,
		},
		{
			name: "draw_four_wild",
			fields: fields{
				Value: draw4WildValue,
			},
			want: undefinedColor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			if got := c.getNextColor(); got != tt.want {
				t.Errorf("card.getNextColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_isSkip(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zero_red",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			want: false,
		},
		{
			name: "draw_two",
			fields: fields{
				Value: draw2Value,
				Color: redColor,
			},
			want: true,
		},
		{
			name: "wild_card",
			fields: fields{
				Value: wildValue,
			},
			want: false,
		},
		{
			name: "draw_four_wild",
			fields: fields{
				Value: draw4WildValue,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			if got := c.isSkip(); got != tt.want {
				t.Errorf("card.isSkip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_isReverse(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zero_red",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			want: false,
		},
		{
			name: "draw_two",
			fields: fields{
				Value: draw2Value,
				Color: redColor,
			},
			want: false,
		},
		{
			name: "wild_card",
			fields: fields{
				Value: wildValue,
			},
			want: false,
		},
		{
			name: "draw_four_wild",
			fields: fields{
				Value: draw4WildValue,
			},
			want: false,
		},
		{
			name: "reverse_yellow",
			fields: fields{
				Value: reverseValue,
				Color: yellowColor,
			},
			want: true,
		},
		{
			name: "reverse_red",
			fields: fields{
				Value: reverseValue,
				Color: redColor,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			if got := c.isReverse(); got != tt.want {
				t.Errorf("card.isReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_isNumber(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zero_red",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			want: true,
		},
		{
			name: "nine_blue",
			fields: fields{
				Value: "9",
				Color: blueColor,
			},
			want: true,
		},
		{
			name: "ten_blue",
			fields: fields{
				Value: "10",
				Color: blueColor,
			},
			want: false,
		},
		{
			name: "draw_two",
			fields: fields{
				Value: draw2Value,
				Color: redColor,
			},
			want: false,
		},
		{
			name: "wild_card",
			fields: fields{
				Value: wildValue,
			},
			want: false,
		},
		{
			name: "draw_four_wild",
			fields: fields{
				Value: draw4WildValue,
			},
			want: false,
		},
		{
			name: "reverse_yellow",
			fields: fields{
				Value: reverseValue,
				Color: yellowColor,
			},
			want: false,
		},
		{
			name: "reverse_red",
			fields: fields{
				Value: reverseValue,
				Color: redColor,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			if got := c.isNumber(); got != tt.want {
				t.Errorf("card.isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_isValid(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zero_red",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			want: true,
		},
		{
			name: "nine_blue",
			fields: fields{
				Value: "9",
				Color: blueColor,
			},
			want: true,
		},
		{
			name: "ten_blue",
			fields: fields{
				Value: "10",
				Color: blueColor,
			},
			want: false,
		},
		{
			name: "zero_purple",
			fields: fields{
				Value: "0",
				Color: "purple",
			},
			want: false,
		},
		{
			name: "invalid_value",
			fields: fields{
				Value: "not_valid",
				Color: redColor,
			},
			want: false,
		},
		{
			name: "draw_two",
			fields: fields{
				Value: draw2Value,
				Color: redColor,
			},
			want: true,
		},
		{
			name: "wild_card",
			fields: fields{
				Value: wildValue,
			},
			want: true,
		},
		{
			name: "draw_four_wild",
			fields: fields{
				Value: draw4WildValue,
			},
			want: true,
		},
		{
			name: "reverse_yellow",
			fields: fields{
				Value: reverseValue,
				Color: yellowColor,
			},
			want: true,
		},
		{
			name: "reverse_red",
			fields: fields{
				Value: reverseValue,
				Color: redColor,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			if got := c.isValid(); got != tt.want {
				t.Errorf("card.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_CanPlay(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	type args struct {
		d            []engine.Card
		currentColor string
	}

	deck := []engine.Card{
		&card{
			Value: "0",
			Color: blueColor,
		},
	}

	deckWild := []engine.Card{
		&card{
			Value: wildValue,
			Color: blackColor,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "can_play_same_color_diff_value",
			fields: fields{
				Value: reverseValue,
				Color: blueColor,
			},
			args: args{
				d: deck,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "can_play_diff_color_same_value",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			args: args{
				d: deck,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "can_play_wild",
			fields: fields{
				Value: wildValue,
				Color: blackColor,
			},
			args: args{
				d: deck,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "can_play_after_wild",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			args: args{
				d:            deckWild,
				currentColor: redColor,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "cannot_play_after_wild",
			fields: fields{
				Value: "0",
				Color: redColor,
			},
			args: args{
				d:            deckWild,
				currentColor: blueColor,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "cannot_play_diff_color_diff_value",
			fields: fields{
				Value: draw2Value,
				Color: redColor,
			},
			args: args{
				d: deck,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			got, err := c.CanPlay(tt.args.currentColor, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("card.CanPlay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("card.CanPlay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_Play(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	type args struct {
		d            []engine.Card
		currentColor string
	}

	deck := []engine.Card{
		&card{
			Value: "0",
			Color: blueColor,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    engine.Result
		wantErr bool
	}{
		{
			name: "play_same_color_diff_value",
			fields: fields{
				Value: reverseValue,
				Color: blueColor,
			},
			args: args{
				d: deck,
			},
			want: &result{
				isReverse: true,
				isSkip:    false,
				drawCount: 0,
				nextColor: blueColor,
				deck:      append(deck, &card{Value: reverseValue, Color: blueColor}),
			},
			wantErr: false,
		},
		{
			name: "play_not_valid_card",
			fields: fields{
				Value: reverseValue,
				Color: redColor,
			},
			args: args{
				d: deck,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "play_wild",
			fields: fields{
				Value: wildValue,
				Color: blackColor,
			},
			args: args{
				d: deck,
			},
			want: &result{
				isReverse: false,
				isSkip:    false,
				drawCount: 0,
				nextColor: undefinedColor,
				deck:      append(deck, &card{Value: wildValue, Color: blackColor}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			got, err := c.Play(tt.args.currentColor, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("card.Play() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("card.Play() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_card_isColorValid(t *testing.T) {
	type fields struct {
		Value string
		Color string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "invalid_color",
			fields: fields{
				Value: "not_valid",
				Color: "not_vaid",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				Value: tt.fields.Value,
				Color: tt.fields.Color,
			}
			if got := c.isColorValid(); got != tt.want {
				t.Errorf("card.isColorValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
