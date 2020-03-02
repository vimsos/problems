package dailycodeproblem

import (
	"reflect"
	"testing"
)

func Test_productOfAllOthers(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{input: []int{1, 2, 3, 4, 5}}, []int{120, 60, 40, 30, 24}},
		{"2", args{input: []int{0, 2, 1}}, []int{2, 0, 0}},
		{"3", args{input: []int{3, -2, 1}}, []int{-2, 3, -6}},
		{"4", args{input: []int{1, 0, 3, 0, 5}}, []int{0, 0, 0, 0, 0}},
		{"5", args{input: []int{55}}, nil},
		{"6", args{input: []int{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productOfAllOthers(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productOfAllOthers() = %v, want %v", got, tt.want)
			}
			if got := leftRightProductOfAllOthers(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftRightProductOfAllOthers() = %v, want %v", got, tt.want)
			}
		})
	}
}
