package leetcode

// Given a binary tree, each node has value 0 or 1.  Each root-to-leaf path represents a binary number starting with the most significant bit.  For example, if the path is
// 0 -> 1 -> 1 -> 0 -> 1,
// then this could represent
// 01101 in binary, which is 13.
// For all leaves in the tree, consider the numbers represented by the path from the root to that leaf.
// Return the sum of these numbers.

type node01 struct {
	val   int
	left  *node01
	right *node01
}

func deserializeTree01(tree []int) *node01 {
	t := append([]int{}, tree...)
	return deserializeNode01(&t)
}

func deserializeNode01(tree *[]int) *node01 {
	if tree == nil || len(*tree) == 0 || (*tree)[0] == -1 {
		return nil
	}

	val := (*tree)[0]
	var left, right *node01
	if len(*tree) > 1 {
		*tree = (*tree)[1:]
		left = deserializeNode01(tree)
	}
	if len(*tree) > 1 {
		*tree = (*tree)[1:]
		right = deserializeNode01(tree)
	}

	return &node01{val: val, left: left, right: right}
}

func traverseSumRootToLeaf(n *node01, v int, sum *int) {
	if n == nil {
		return
	}
	v = v*2 + n.val

	if n.left == nil && n.right == nil {
		*sum += v
		return
	}
	traverseSumRootToLeaf(n.left, v, sum)
	traverseSumRootToLeaf(n.right, v, sum)
}

func sumRootToLeaf(root *node01) int {
	var sum int
	traverseSumRootToLeaf(root, 0, &sum)
	return sum
}
