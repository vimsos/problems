package dcp

// The power set of a set is the set of all its subsets. Write a function that, given a set, generates its power set.
// For example, given the set {1, 2, 3}, it should return {{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}.
// You may also use a list or array to represent a set.

func powerSet(set []int) [][]int {
	checkBit := func(n int, pos uint) bool {
		val := n & (1 << pos)
		return (val > 0)
	}

	var result [][]int
	total := 1 << len(set)
	for i := 0; i < total; i++ {
		var sub []int
		for j := 0; j < len(set); j++ {
			if checkBit(i, uint(j)) {
				sub = append(sub, set[j])
			}
		}
		result = append(result, sub)
	}

	return result
}
