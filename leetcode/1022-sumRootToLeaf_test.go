package leetcode

import "testing"

func Test_sumRootToLeaf(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{tree: []int{1, 0, 0, -1, -1, 1, -1, -1, 1, 0, 1}}, 22},
		{"1", args{tree: nil}, 0},
		{"2", args{tree: []int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumRootToLeaf(deserializeTree01(tt.args.tree)); got != tt.want {
				t.Errorf("sumRootToLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
