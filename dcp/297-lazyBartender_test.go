package dcp

import "testing"

func Test_lazyBartender(t *testing.T) {
	type args struct {
		favorites map[int][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{favorites: map[int][]int{0: {0, 1, 3, 6}, 1: {1, 4, 7}, 2: {2, 4, 7, 5}, 3: {3, 2, 5}, 4: {5, 8}}}, 2},
		{"1", args{favorites: map[int][]int{0: {0}, 1: {1, 4, 7}, 2: {2, 5}, 3: {3, 5}, 4: {5, 8}}}, 3},
		{"2", args{favorites: map[int][]int{0: {5}, 1: {5}, 2: {7, 5}, 3: {3, 2, 5}, 4: {5, 8}}}, 1},
		{"3", args{favorites: map[int][]int{0: {0}, 1: {1}, 2: {2}, 3: {3}, 4: {4}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lazyBartender(tt.args.favorites); got != tt.want {
				t.Errorf("lazyBartender() = %v, want %v", got, tt.want)
			}
		})
	}
}
