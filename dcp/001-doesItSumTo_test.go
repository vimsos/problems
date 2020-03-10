package dcp

import "testing"

func TestDoesItSumTo(t *testing.T) {
	type args struct {
		k   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{k: 17, arr: []int{10, 15, 3, 7}}, true},
		{"2", args{k: 55, arr: []int{10, 15, 3, 7}}, false},
		{"3", args{k: -12, arr: []int{-8, -6, -4, 4}}, true},
		{"4", args{k: -34, arr: []int{0, 0, 77, 12}}, false},
		{"5", args{k: 0, arr: []int{0, 0, 77, 12}}, true},
		{"6", args{k: 0, arr: []int{0, 77, 12}}, false},
		{"7", args{k: 0, arr: []int{-128, 128, 77, 12}}, true},
		{"8", args{k: 0, arr: []int{-128}}, false},
		{"9", args{k: 0, arr: []int{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doesItSumTo(tt.args.k, tt.args.arr); got != tt.want {
				t.Errorf("doesItSumTo() = %v, want %v", got, tt.want)
			}
			if got := mapDoesItSumTo(tt.args.k, tt.args.arr); got != tt.want {
				t.Errorf("mapDoesItSumTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
