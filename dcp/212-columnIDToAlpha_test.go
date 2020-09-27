package dcp

import "testing"

func Test_columnIDToAlpha(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{id: 1}, "A"},
		{"1", args{id: 27}, "AA"},
		{"2", args{id: 10}, "J"},
		{"3", args{id: 53}, "BA"},
		{"4", args{id: 26}, "Z"},
		{"5", args{id: 52}, "AZ"},
		{"6", args{id: 701}, "ZY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := columnIDToAlpha(tt.args.id); got != tt.want {
				t.Errorf("columnIDToAlpha(%v) = %v, want %v",tt.args.id, got, tt.want)
			}
		})
	}
}
