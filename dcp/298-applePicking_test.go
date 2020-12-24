package dcp

import "testing"

func Test_applePicking(t *testing.T) {
	type args struct {
		types []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{types: []int{2, 1, 2, 3, 3, 1, 3, 5}}, 4},
		{"1", args{types: []int{}}, 0},
		{"2", args{types: []int{2, 1, 2, 3, 1, 3, 5}}, 3},
		{"3", args{types: []int{1, 2, 3, 4, 1, 2, 3, 4}}, 2},
		{"4", args{types: []int{2, 1, 2, 2, 2, 1, 2, 1}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applePicking(tt.args.types); got != tt.want {
				t.Errorf("applePicking() = %v, want %v", got, tt.want)
			}
		})
	}
}
