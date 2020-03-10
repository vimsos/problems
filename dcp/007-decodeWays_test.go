package dcp

import "testing"

func Test_waysToDecode(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{msg: "111453"}, 5},
		{"1", args{msg: "111"}, 3},
		{"2", args{msg: "12"}, 2},
		{"3", args{msg: "126"}, 3},
		{"4", args{msg: "10"}, 1},
		{"5", args{msg: "1787897759966261825913315262377298132516969578441236833255596967132573482281598412163216914566534565"}, 5898240},
	}
	for _, tt := range tests {
		// t.Run(tt.name, func(t *testing.T) {
		// 	if got := decodeWays(tt.args.msg); got != tt.want {
		// 		t.Errorf("decodeWays() = %v, want %v", got, tt.want)
		// 	}
		// })
		t.Run(tt.name, func(t *testing.T) {
			if got := memDecodeWays(tt.args.msg); got != tt.want {
				t.Errorf("memDecodeWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
