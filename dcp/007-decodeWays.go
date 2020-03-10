package dcp

import (
	"strconv"
)

// Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.
// For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.
// You can assume that the messages are decodable. For example, '001' is not allowed.

//from https://leetcode.com/problems/decode-ways/discuss/30451/Evolve-from-recursion-to-dp
func memDecodeWays(s string) int {
	runes := []rune(s)
	len := len(runes)
	mem := make([]int, len+1)
	mem[len] = 1
	for i := 0; i < len; i++ {
		mem[i] = -1
	}
	return memCountWays(0, &runes, &mem)
}

func memCountWays(p int, r *[]rune, m *[]int) int {
	if (*m)[p] > -1 {
		return (*m)[p]
	}

	len := len(*r)
	if (*r)[p] == '0' {
		(*m)[p] = 0
		return 0
	}

	ways := memCountWays(p+1, r, m)

	if len-p > 1 {
		if val, err := strconv.Atoi(string(*r)[p : p+2]); err == nil && val > 1 && val < 27 {
			ways += memCountWays(p+2, r, m)
		}
	}

	(*m)[p] = ways
	return ways
}

func decodeWays(msg string) int {
	runes := []rune(msg)
	return countWays(&runes)
}

func countWays(r *[]rune) int {
	len := len(*r)

	if len == 0 {
		return 1
	}
	if (*r)[0] == '0' {
		return 0
	}

	minusOne := (*r)[1:]
	ways := countWays(&minusOne)

	if len > 1 {
		if val, err := strconv.Atoi(string(*r)[0:2]); err == nil && val > 1 && val < 27 {
			minusTwo := (*r)[2:]
			ways += countWays(&minusTwo)
		}
	}
	return ways
}
