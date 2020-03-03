package dailycodeproblem

import "testing"

func Test_firstMissingPositiveInteger(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{arr: []int{3, 4, -1, 1}}, 2},
		{"2", args{arr: []int{1, 2, 0}}, 3},
		{"3", args{arr: []int{2, 2, 0}}, 1},
		{"4", args{arr: []int{12, 8, -3, 1}}, 2},
		{"5", args{arr: []int{-1, 3, 1, 4}}, 2},
		{"6", args{arr: []int{4, 3, 2, 1}}, 5},
		{"7", args{arr: []int{4, 2, 3, 1}}, 5},
		{"8", args{arr: []int{-4, -3, -2, 1}}, 2},
		{"9", args{arr: []int{44, 33, 3, 1}}, 2},
		{"10", args{arr: []int{5, 4, -3, -2, 1}}, 2},
		{"11", args{arr: []int{6, 5, 4, -3, -2, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapFirstMissingPositiveInteger(tt.args.arr); got != tt.want {
				t.Errorf("mapFirstMissingPositiveInteger() = %v, want %v", got, tt.want)
			}
			if got := firstMissingPositiveInteger(tt.args.arr); got != tt.want {
				t.Errorf("firstMissingPositiveInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
