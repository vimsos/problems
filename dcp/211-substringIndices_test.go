package dcp

import (
	"reflect"
	"testing"
)

func Test_substringIndices(t *testing.T) {
	type args struct {
		str string
		ptn string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{str: "abracadabra", ptn: "abr"}, []int{0, 7}},
		{"1", args{str: "abracadabra", ptn: "abc"}, []int{}},
		{"2", args{str: "the book is on the table", ptn: "the "}, []int{0, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substringIndices(tt.args.str, tt.args.ptn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("substringIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
