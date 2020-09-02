package dcp

import (
	"reflect"
	"testing"
)

func Test_permutations(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"0", args{arr: []int{1, 2, 3}}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"1", args{arr: []int{1, 2, 3, 4}}, [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 2, 3}, {1, 4, 3, 2}, {2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 1, 3}, {2, 4, 3, 1}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 4, 1, 2}, {3, 4, 2, 1}, {4, 1, 2, 3}, {4, 1, 3, 2}, {4, 2, 1, 3}, {4, 2, 3, 1}, {4, 3, 1, 2}, {4, 3, 2, 1}}},
		{"2", args{arr: []int{8, 9}}, [][]int{{8, 9}, {9, 8}}},
	}
	for _, tt := range tests {
		arrayContains := func(arr [][]int, item []int) bool {
			for _, i := range arr {
				if reflect.DeepEqual(i, item) {
					return true
				}
			}
			return false
		}

		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.arr)

			if len(got) != len(tt.want) {
				t.Errorf("len(permutations()) = %v, want %v", len(got), len(tt.want))
				return
			}

			for _, i := range tt.want {
				if !arrayContains(got, i) {
					t.Errorf("permutations() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}

}
