package dcp

import (
	"reflect"
	"testing"
)

func Test_regularNumbers(t *testing.T) {
	numbers := []uint64{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 25, 27, 30, 32, 36, 40, 45, 48, 50, 54, 60, 64, 72, 75, 80, 81, 90, 96, 100, 108, 120, 125, 128, 135, 144, 150, 160, 162, 180, 192, 200, 216, 225, 240, 243, 250, 256, 270, 288, 300, 320, 324, 360, 375, 384, 400, 405}

	type args struct {
		n uint64
	}
	tests := []struct {
		name        string
		args        args
		wantNumbers []uint64
	}{
		{"0", args{n: 0}, []uint64{}},
		{"1", args{n: 5}, numbers[:5]},
		{"2", args{n: 13}, numbers[:13]},
		{"3", args{n: 26}, numbers[:26]},
		{"4", args{n: 62}, numbers[:62]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumbers := regularNumbers(tt.args.n); !reflect.DeepEqual(gotNumbers, tt.wantNumbers) {
				t.Errorf("regularNumbers(%v) = %v, want %v", tt.args.n, gotNumbers, tt.wantNumbers)
			}
		})
	}
}
