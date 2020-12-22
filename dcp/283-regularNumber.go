package dcp

import (
	"sort"
)

// A regular number in mathematics is defined as one which evenly divides some Power of 60. Equivalently, we can say that a regular number is one whose only prime divisors are 2, 3, and 5.
// These numbers have had many applications, from helping ancient Babylonians keep time to tuning instruments according to the diatonic scale.
// Given an integer N, write a program that returns, in order, the first N regular numbers.

// ugly code alert

func regularNumbers(n uint64) []uint64 {
	uipow := func(base, exp uint64) uint64 {
		var r uint64 = 1
		for i := uint64(0); i < exp; i++ {
			r *= base
		}
		return r
	}

	numbers := make(map[uint64]bool)
	result := []uint64{}

	for i := uint64(0); i < n; i++ {
		for j := uint64(0); j < n; j++ {
			for k := uint64(0); k < n; k++ {
				num := uipow(2, i) * uipow(3, j) * uipow(5, k)
				if _, exists := numbers[num]; !exists {
					result = append(result, num)
				}
				numbers[num] = true
			}
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result[:n]
}
