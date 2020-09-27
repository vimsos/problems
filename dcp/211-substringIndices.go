package dcp

// Given a string and a pattern, find the starting indices of all occurrences of the pattern in the string. For example, given the string "abracadabra" and the pattern "abr", you should return [0, 7].

func substringIndices(str string, ptn string) []int {
	indices := make([]int, 0)
	l := len(ptn)

	for i := 0; i+l < len(str); i++ {
		if str[i:i+l] == ptn {
			indices = append(indices, i)
		}
	}

	return indices
}
