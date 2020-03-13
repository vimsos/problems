package dcp

import (
	"math"
	"testing"
)

func Test_mcpi(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{"1", math.Pi},
		{"1", math.Pi},
		{"1", math.Pi},
		{"1", math.Pi},
		{"1", math.Pi},
		{"1", math.Pi},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mcpi(); math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("mcpi() = %v, want %v", got, tt.want)
			}
		})
	}
}
