package dcp

import (
	"reflect"
	"testing"
)

func Test_waterCost(t *testing.T) {
	type args struct {
		pipeSet map[string]map[string]uint
	}
	tests := []struct {
		name string
		args args
		want map[string]map[string]uint
	}{
		{"0",
			args{map[string]map[string]uint{
				"plant": {"A": 1, "B": 5, "C": 20},
				"A":     {"C": 15},
				"B":     {"C": 10},
				"C":     {}},
			},
			map[string]map[string]uint{"plant": {"A": 1, "B": 5}, "B": {"C": 10}},
		},
		{"1",
			args{map[string]map[string]uint{
				"plant": {"A": 1, "B": 1, "C": 1},
				"A":     {"C": 15, "B": 25},
				"B":     {"C": 10},
				"C":     {}},
			},
			map[string]map[string]uint{"plant": {"A": 1, "B": 1, "C": 1}},
		},
		{"2",
			args{map[string]map[string]uint{
				"plant": {"A": 25, "B": 15},
				"A":     {"C": 5},
				"B":     {"C": 7, "A": 8},
				"C":     {}},
			},
			map[string]map[string]uint{"plant": {"B": 15}, "B": {"A": 8}, "A": {"C": 5}},
		},
		{"3",
			args{map[string]map[string]uint{
				"plant": {"A": 25, "B": 15, "T": 60, "Y": 30},
				"A":     {"C": 5, "T": 40},
				"B":     {"C": 7, "A": 8},
				"C":     {"Y": 20, "T": 100},
				"Y":     {"T": 10},
				"T":     {"B": 80}},
			},
			map[string]map[string]uint{"plant": {"B": 15}, "B": {"A": 8}, "A": {"C": 5}, "C": {"Y": 20}, "Y": {"T": 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waterCost(tt.args.pipeSet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("waterCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
