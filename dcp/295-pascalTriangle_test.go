package dcp

import (
	"reflect"
	"testing"
)

func Test_pascalTriangle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{n: 0}, []int{1}},
		{"1", args{n: 1}, []int{1, 1}},
		{"2", args{n: 2}, []int{1, 2, 1}},
		{"3", args{n: 3}, []int{1, 3, 3, 1}},
		{"4", args{n: 4}, []int{1, 4, 6, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pascalTriangle(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pascalTriangle(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
