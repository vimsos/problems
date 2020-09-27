package dcp

// find the starting indices of all occurrences of the patterne

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
