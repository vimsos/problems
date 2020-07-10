package dcp

import (
	"reflect"
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
		{"1", args{input: []string{" dog", "  deer  ", "deal"}, query: "de"}, []string{"deal", "deer"}},
		{"2", args{input: []string{"edison", "fluxo laminar"}, query: "e"}, []string{"edison"}},
		{"3", args{input: []string{"amélia"}, query: "amé"}, []string{"amélia"}},
		{"4", args{input: []string{"fluxo laminar"}, query: "f"}, []string{"fluxo laminar"}},
		{"5", args{input: []string{" dog", "  deer  ", "deal"}, query: "d"}, []string{"dog", "deal", "deer"}},
	}

	arrayContains := func(arr []string, item string) bool {
		// this is also ugly code
		for _, i := range arr {
			if reflect.DeepEqual(i, item) || len(i) == 0 && len(item) == 0 {
				return true
			}
		}
		return false
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &trieNode{}
			root.initTrieRoot('*')
			root.insertMany(tt.args.input)

			got := root.query(tt.args.query)

			if len(got) != len(tt.want) {
				t.Errorf("len(powerSet()) = %v, want %v", len(got), len(tt.want))
			}

			for _, i := range tt.want {
				if !arrayContains(got, i) {
					t.Errorf("powerSet() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
