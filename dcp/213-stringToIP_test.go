package dcp

import (
	"reflect"
	"testing"
)

func Test_stringToIP(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"0", args{str: "2542540123"}, []string{"254.25.40.123", "254.254.0.123"}},
		{"1", args{str: "10001"}, []string{"10.0.0.1"}},
		{"2", args{str: "1111"}, []string{"1.1.1.1"}},
		{"3", args{str: "21658202238"}, []string{"216.58.202.238"}},
		{"4", args{str: "17221729110"}, []string{"172.217.29.110"}},
		{"5", args{str: "16161616"}, []string{"16.16.16.16", "16.16.161.6", "16.161.6.16","16.161.61.6","161.6.16.16","161.6.161.6","161.61.6.16","161.61.61.6"}},
		{"6", args{str: "172217291100"}, []string{}},
		{"7", args{str: "17221729110aa"}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToIP(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
