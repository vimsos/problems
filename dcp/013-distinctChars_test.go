package dcp

import "testing"

func Test_distinctChars(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{s: "abcba", k: 2}, 3},
		{"2", args{s: "eceba", k: 2}, 3},
		{"3", args{s: "ecebbbba", k: 2}, 5},
		{"4", args{s: "ecebbbba", k: 3}, 7},
		{"5", args{s: "j", k: 5}, 1},
		{"6", args{s: "", k: 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctChars(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("distinctChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
