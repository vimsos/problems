package dcp

// Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.
// For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.
// You can modify the input array in-place.

//from https://stackoverflow.com/a/51346319
func firstMissingPositiveInteger(arr []int) int {
	//function output is bound between 1 and len(arr)+1
	abs := func(n int) int {
		if n < 0 {
			n = -n
		}
		return n
	}

	end := len(arr)

	//sort negative values to the end of the array
	//end becomes the lenght of the positive side of the array
	for i := 0; i < end; i++ {
		if arr[i] < 1 {
			arr[end-1], arr[i] = arr[i], arr[end-1]
			end--
			i--
		}
	}

	//iterate from 0 to end
	//flag the index of numbers that are smaller or equal to end as present using the negative sign
	for i := 0; i < end; i++ {
		val := abs(arr[i])
		if val <= end {
			//use abs to avoid unflagging already flagged index
			arr[val-1] = -abs(arr[val-1])
		}
	}

	//iterate from 0 to end again
	//if a position contains a positive number that indicates the index+1 of that position isn't present in the array
	for i := 0; i < end; i++ {
		if arr[i] > 0 {
			return i + 1
		}
	}

	//in case there are no positive numbers from 0 to end, it indicates all numbers between 1 and end are present in the array, the answer will be end+1
	return end + 1
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
