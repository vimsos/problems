package dcp

import (
	"fmt"
	"testing"
)

func Test_insertManyQuery(t *testing.T) {
	type args struct {
		input []string
		query string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"0", args{input: []string{}, query: "e"}, []string{}},
		{"1", args{input: []string{" dog", "  deer  ", "deal"}, query: "de"}, []string{"deer", "deal"}},
		{"2", args{input: []string{"edison", "fluxo laminar"}, query: "e"}, []string{"edison"}},
		{"3", args{input: []string{"amélia"}, query: "amé"}, []string{"amélia"}},
		{"4", args{input: []string{""}, query: "f"}, []string{"fluxo laminar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertMany(tt.args.input)

			if got := query(tt.args.query); fmt.Sprint(got) != fmt.Sprint(tt.want) {
				t.Errorf("query(%q) = %v, want %v", tt.args.query, got, tt.want)
			}
		})
	}
}
