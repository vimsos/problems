package dcp

import (
	"reflect"
	"testing"
)

func Test_maxValueSubArray(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{arr: []int{10, 5, 2, 7, 8, 7}, k: 3}, []int{10, 7, 8, 8}},
		{"2", args{arr: []int{1, 2, 3, 1, 4, 5, 2, 3, 6}, k: 3}, []int{3, 3, 4, 5, 5, 5, 6}},
		{"3", args{arr: []int{8, 5, 10, 7, 9, 4, 15, 12, 90, 13}, k: 4}, []int{10, 10, 10, 15, 15, 90, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValueSubArray(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxValueSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dpMaxValueSubArray(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{arr: []int{10, 5, 2, 7, 8, 7}, k: 3}, []int{10, 7, 8, 8}},
		{"2", args{arr: []int{1, 2, 3, 1, 4, 5, 2, 3, 6}, k: 3}, []int{3, 3, 4, 5, 5, 5, 6}},
		{"3", args{arr: []int{8, 5, 10, 7, 9, 4, 15, 12, 90, 13}, k: 4}, []int{10, 10, 10, 15, 15, 90, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deqMaxValueSubArray(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deqMaxValueSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
