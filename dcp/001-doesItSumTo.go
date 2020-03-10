package dcp

// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
// Bonus: Can you do this in one pass?

func doesItSumTo(k int, arr []int) bool {
	for ii, vi := range arr {
		for ij, vj := range arr {
			if ii == ij {
				continue
			}

			if vi+vj == k {
				return true
			}
		}
	}
	return false
}

func mapDoesItSumTo(k int, arr []int) bool {
	m := make(map[int]bool)

	for _, v := range arr {
		need := k - v

		if m[need] == true {
			return true
		}

		m[v] = true
	}
	return false
}
