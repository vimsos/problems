package dcp

import "testing"

func Test_longestConsecutive1s(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{n: 156}, 3},
		{"1", args{n: 0}, 0},
		{"2", args{n: 1}, 1},
		{"3", args{n: 2}, 1},
		{"4", args{n: 3}, 2},
		{"5", args{n: 5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive1s(tt.args.n); got != tt.want {
				t.Errorf("longestConsecutive1s() = %v, want %v", got, tt.want)
			}
		})
	}
}
