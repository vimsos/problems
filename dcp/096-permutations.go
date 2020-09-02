package dcp

// Given a number in the form of a list of digits, return all possible permutations.
// For example, given [1,2,3], return [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]].

func permute(set []int) [][]int {
	var result [][]int
	permutationsHelper(set, []int{}, &result)
	return result
}

func permutationsHelper(set []int, sequence []int, permutations *[][]int) {
	removeAtIndex := func(arr []int, index int) []int {
		var result []int
		result = append(result, arr[:index]...)
		result = append(result, arr[index+1:]...)
		return result
	}

	if len(set) == 0 {
		*permutations = append(*permutations, sequence)
		return
	}

	for i := 0; i < len(set); i++ {
		nextSequence := append(sequence)
		nextSequence = append(sequence, set[i])

		nextSet := append(set)
		nextSet = removeAtIndex(nextSet, i)

		permutationsHelper(nextSet, nextSequence, permutations)
	}
}
