package dailycodeproblem

import (
	"fmt"
	"strings"
)

// Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.
// For example, given the following Node class

// class Node:
//     def __init__(self, val, left=None, right=None):
//         self.val = val
//         self.left = left
//         self.right = right

// The following test should pass:

// node = Node('root', Node('left', Node('left.left')), Node('right'))
// assert deserialize(serialize(node)).left.left.val == 'left.left'

type node struct {
	val   string
	left  *node
	right *node
}

func serializeNode(n *node, s *[]string) {
	if n != nil {
		*s = append(*s, n.val)
		serializeNode(n.left, s)
		serializeNode(n.right, s)
		return
	}

	*s = append(*s, "nil")
}

func serialize(root *node) string {
	nodes := make([]string, 0)

	serializeNode(root, &nodes)

	return fmt.Sprint(nodes)
}

func deserializeNodes(s *[]string) *node {
	if len(*s) < 1 {
		return nil
	}

	val := (*s)[0]
	if val == "nil" {
		return nil
	}

	*s = (*s)[1:]
	left := deserializeNodes(s)
	*s = (*s)[1:]
	right := deserializeNodes(s)

	return &node{val: val, left: left, right: right}
}

func deserialize(input string) *node {
	input = strings.ReplaceAll(input, "[", "")
	input = strings.ReplaceAll(input, "]", "")
	split := strings.Split(input, " ")

	return deserializeNodes(&split)
}
