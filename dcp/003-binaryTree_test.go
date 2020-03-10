package dcp

import (
	"testing"
)

func Test_serdes(t *testing.T) {
	original := &node{val: "root", left: &node{val: "left", left: &node{val: "left.left"}}, right: &node{val: "right"}}

	reconstructed := deserialize(serialize(original))

	if original.val != reconstructed.val {
		t.Errorf("serialize(deserialize()) = %v, want %v", original.val, reconstructed.val)
	}
	if original.left.val != reconstructed.left.val {
		t.Errorf("serialize(deserialize()) = %v, want %v", original.left.val, reconstructed.left.val)
	}
	if original.left.left.val != reconstructed.left.left.val {
		t.Errorf("serialize(deserialize()) = %v, want %v", original.left.left.val, reconstructed.left.left.val)
	}
	if original.right.val != reconstructed.right.val {
		t.Errorf("serialize(deserialize()) = %v, want %v", original.right.val, reconstructed.right.val)
	}
	if original.right.left != reconstructed.right.left {
		t.Errorf("serialize(deserialize()) = %v, want %v", original.val, reconstructed.val)
	}
	if serialize(original) != serialize(reconstructed) {
		t.Errorf("serialize(deserialize()) = %v, want %v", serialize(original), serialize(reconstructed))
	}
}

func Test_serialize(t *testing.T) {
	type args struct {
		root *node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{}, "[nil]"},
		{"2", args{root: &node{val: "root", left: &node{val: "left", left: &node{val: "left.left"}}, right: &node{val: "right"}}}, "[root left left.left nil nil nil right nil nil]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
