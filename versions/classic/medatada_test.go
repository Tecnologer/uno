package main

import "testing"

func Test_metadata_GetMaxPlayer(t *testing.T) {
	tests := []struct {
		name string
		me   metadata
		want int
	}{
		{
			name: "GetMaxPlayer",
			me:   0,
			want: 6,
		},
		{
			name: "GetMaxPlayer",
			me:   1,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.me.GetMaxPlayer(); got != tt.want {
				t.Errorf("GetMaxPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metadata_GetMinPlayer(t *testing.T) {
	tests := []struct {
		name string
		me   metadata
		want int
	}{
		{
			name: "GetMinPlayer",
			me:   0,
			want: 2,
		},
		{
			name: "GetMinPlayer",
			me:   1,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.me.GetMinPlayer(); got != tt.want {
				t.Errorf("GetMinPlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metadata_GetName(t *testing.T) {
	tests := []struct {
		name string
		me   metadata
		want string
	}{
		{
			name: "GetName",
			me:   0,
			want: "Classic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.me.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metadata_GetVersion(t *testing.T) {
	tests := []struct {
		name string
		me   metadata
		want string
	}{
		{
			name: "GetVersion",
			me:   0,
			want: "0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.me.GetVersion(); got != tt.want {
				t.Errorf("GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metadata_String(t *testing.T) {
	tests := []struct {
		name string
		m    metadata
		want string
	}{
		{
			name: "String",
			m:    0,
			want: "Classic [players: 2-6] (version: 0.0)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
