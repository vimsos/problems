package dcp

import "testing"

func Test_romanToDecimal(t *testing.T) {
	type args struct {
		roman string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{roman: "XIV"}, 14},
		{"1", args{roman: ""}, 0},
		{"2", args{roman: "IV"}, 4},
		{"3", args{roman: "V"}, 5},
		{"4", args{roman: "XIII"}, 13},
		{"5", args{roman: "MDCLXVI"}, 1666},
		{"6", args{roman: "CM"}, 900},
		{"7", args{roman: "CCLIV"}, 254},
		{"8", args{roman: "LLL"}, 150},
		{"9", args{roman: "X"}, 10},
		{"10", args{roman: "I"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToDecimal(tt.args.roman); got != tt.want {
				t.Errorf("romanToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
