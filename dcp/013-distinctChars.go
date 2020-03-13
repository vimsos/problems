package dcp

// Given an integer k and a string s, find the length of the longest substring that contains at most k distinct characters.
// For example, given s = "abcba" and k = 2, the longest substring with k distinct characters is "bcb".

func distinctChars(s string, k int) int {
	distinct := func(r []rune) int {
		m := make(map[rune]int)
		for _, v := range r {
			m[v]++
		}
		return len(m)
	}

	r := []rune(s)
	var longest int

	for i := 0; i < len(r); i++ {
		for j := len(r); j > i && j-i >= longest; j-- {
			if distinct(r[i:j]) <= k && j-i > longest {
				longest = j - i
			}
		}
	}

	return longest
}
