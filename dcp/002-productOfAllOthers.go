package dcp

// Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.
// For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].
// Follow-up: what if you can't use division?

func productOfAllOthers(input []int) []int {
	len := len(input)
	if len < 2 {
		return nil
	}

	output := make([]int, len)
	for i := 0; i < len; i++ {
		output[i] = 1
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if i == j {
				continue
			}
			output[j] *= input[i]
		}
	}

	return output
}

func leftRightProductOfAllOthers(input []int) []int {
	len := len(input)
	if len < 2 {
		return nil
	}

	output := make([]int, len)
	left := make([]int, len)
	right := make([]int, len)
	left[0] = 1
	right[len-1] = 1

	for i := 1; i < len; i++ {
		left[i] = left[i-1] * input[i-1]
	}
	for i := len - 2; i >= 0; i-- {
		right[i] = right[i+1] * input[i+1]
	}
	for i := 0; i < len; i++ {
		output[i] = left[i] * right[i]
	}

	return output
}
