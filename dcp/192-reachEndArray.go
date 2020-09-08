package dcp

// You are given an array of nonnegative integers. Let's say you start at the beginning of the array and are trying to advance to the end. You can advance at most, the number of steps that you're currently on. Determine whether you can get to the end of the array.
// For example, given the array
// [1, 3, 1, 2, 0, 1],
// we can go from indices
// 0 -> 1 -> 3 -> 5
// , so return true.
// Given the array
// [1, 2, 1, 0, 0],
// we can't reach the end, so return false.

func canReachEnd(a []int) bool {
	return tryReachEnd(a, 0)
}

func tryReachEnd(a []int, i int) bool {
	if a[i] == 0 {
		return false
	}
	if a[i]+i > len(a)-1 {
		return true
	}
	reach := false
	for n := a[i]; n > 0; n-- {
		reach = reach || tryReachEnd(a, i+n)
	}
	return reach
}
