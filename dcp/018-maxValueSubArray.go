package dcp

import (
	"math"
)

// Given an array of integers and a number k, where 1 <= k <= length of the array, compute the maximum values of each subarray of length k.
// For example, given array = [10, 5, 2, 7, 8, 7] and k = 3, we should get: [10, 7, 8, 8], since:

//     10 = max(10, 5, 2)
//     7 = max(5, 2, 7)
//     8 = max(2, 7, 8)
//     8 = max(7, 8, 7)

// Do this in O(n) time and O(k) space. You can modify the input array in-place and you do not need to store the results. You can simply print them out as you compute them.

// from https://www.geeksforgeeks.org/sliding-window-maximum-maximum-of-all-subarrays-of-size-k/
func deqMaxValueSubArray(arr []int, k int) []int {
	end := len(arr) - k + 1
	deq := []int{}

	for i := 0; i < k; i++ {
		for len(deq) > 0 && arr[i] > arr[deq[0]] {
			deq = deq[1:]
		}
		deq = append(deq, i)
	}

	for i := k; i < len(arr); i++ {
		arr[i-k] = arr[deq[0]]

		for len(deq) > 0 && deq[0] <= i-k {
			deq = deq[1:]
		}
		for len(deq) > 0 && arr[i] > arr[deq[0]] {
			deq = deq[1:]
		}
		deq = append(deq, i)
	}

	arr[len(arr)-k] = arr[deq[0]]

	return arr[:end]
}

func maxValueSubArray(arr []int, k int) []int {
	max := func(nums ...int) int {
		big := math.MinInt64
		for _, v := range nums {
			if v > big {
				big = v
			}
		}
		return big
	}

	end := len(arr) - k + 1
	for i := 0; i < end; i++ {
		arr[i] = max(arr[i : i+k]...)
	}

	return arr[:end]
}
