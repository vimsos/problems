package dcp

import "testing"

func Test_canReachEnd(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{[]int{1, 3, 1, 2, 0, 1}}, true},
		{"1", args{[]int{1, 2, 1, 0, 0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReachEnd(tt.args.a); got != tt.want {
				t.Errorf("canReachEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
