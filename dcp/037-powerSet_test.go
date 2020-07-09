package dcp

import (
	"reflect"
	"testing"
)

func Test_powerSet(t *testing.T) {
	type args struct {
		set []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"0", args{set: []int{1, 2, 3}}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{"1", args{set: []int{1, 2, 3, 4}}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}, {4}, {1, 4}, {2, 4}, {1, 2, 4}, {3, 4}, {1, 3, 4}, {2, 3, 4}, {1, 2, 3, 4}}},
	}

	arrayContains := func(arr [][]int, item []int) bool {
		// this is some ugly code
		for _, i := range arr {
			if reflect.DeepEqual(i, item) || len(i) == 0 && len(item) == 0 {
				return true
			}
		}
		return false
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := powerSet(tt.args.set)

			if len(got) != len(tt.want) {
				t.Errorf("len(powerSet()) = %v, want %v", len(got), len(tt.want))
			}

			for _, i := range tt.want {
				if !arrayContains(got, i) {
					t.Errorf("powerSet() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
