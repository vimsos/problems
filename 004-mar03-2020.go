package dailycodeproblem

// Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.
// For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.
// You can modify the input array in-place.

func firstMissingPositiveInteger(arr []int) int {
	max := len(arr) + 1

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 && arr[i] < max {
			arr[arr[i]-1], arr[i] = arr[i], arr[arr[i]-1]
		}
	}

	missing := 1
	for _, v := range arr {
		if missing == v {
			missing++
		}
	}

	return missing
}

func mapFirstMissingPositiveInteger(arr []int) int {
	max := len(arr) + 1
	found := make(map[int]bool, max)

	for _, v := range arr {
		if v > 0 && v < max {
			found[v] = true
		}
	}

	for i := 1; i < max; i++ {
		if !found[i] {
			return i
		}
	}

	return max
}
