package dcp

import "testing"

func Test_reverseIntegerBits(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"0", args{n: 0b1110000111100001111000011110000}, 0b0001111000011110000111100001111},
		{"1", args{n: 0b1}, 0b1111111111111111111111111111110},
		{"2", args{n: 0b1111111111111111111111111111110}, 0b1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseIntegerBits(tt.args.n); got != tt.want {
				t.Errorf("reverseIntegerBits(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
