package dcp

// Given an integer n, return the length of the longest consecutive run of 1s in its binary representation.
// For example, given 156, you should return 3.

func longestConsecutive1s(n int) int {
	longest, current := 0, 0

	for n > 0 {
		r := n % 2
		n /= 2
		if r == 0 {
			current = 0
		} else {
			current++
		}
		if current > longest {
			longest = current
		}
	}

	return longest
}
