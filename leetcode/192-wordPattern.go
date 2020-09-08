package leetcode

import (
	"strings"
)

// Given a pattern and a string str, find if str follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.
// Input: pattern = "abba", str = "dog cat cat dog"
// Output: true
// Input:pattern = "abba", str = "dog cat cat fish"
// Output: false
// You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.

func wordPattern(pattern string, str string) bool {
	rMap := make(map[rune]*string)
	wMap := make(map[string]rune)
	runes := []rune(pattern)
	words := strings.Split(str, " ")

	if len(runes) != len(words) {
		return false
	}
	for i, r := range runes {
		if rMap[r] == nil {
			rMap[r] = &words[i]
			wMap[words[i]] = r
		} else if *rMap[r] != words[i] || wMap[*rMap[r]] != r {
			return false
		}
	}
	return true
}
