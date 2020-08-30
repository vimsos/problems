package dcp

import "math"

// Given a 32-bit integer, return the number with its bits reversed.

// For example, given the binary number
// 1111 0000 1111 0000 1111 0000 1111 0000,
// return
// 0000 1111 0000 1111 0000 1111 0000 1111.

// inverts the bits using the XOR operator, 0^1=1, 1^1=0
func reverseIntegerBits(n int32) int32 {
	return math.MaxInt32 ^ n
}
