package dcp

import "testing"

func Test_lookAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{n: 1}, "1"},
		{"1", args{n: 2}, "11"},
		{"2", args{n: 3}, "21"},
		{"3", args{n: 4}, "1211"},
		{"4", args{n: 5}, "111221"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lookAndSay(tt.args.n); got != tt.want {
				t.Errorf("lookAndSay(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
