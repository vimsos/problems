package dcp

import "testing"

func Test_throwDice(t *testing.T) {
	type args struct {
		n     int
		faces int
		total int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{n: 3, faces: 6, total: 7}, 15},
		{"1", args{n: 1, faces: 2, total: 3}, 0},
		{"2", args{n: 2, faces: 5, total: 10}, 1},
		{"3", args{n: 2, faces: 6, total: 7}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := throwDice(tt.args.n, tt.args.faces, tt.args.total); got != tt.want {
				t.Errorf("throwDice() = %v, want %v", got, tt.want)
			}
		})
	}
}
