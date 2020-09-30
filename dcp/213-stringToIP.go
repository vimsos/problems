package dcp

import "strconv"

// Given a string of digits, generate all possible valid IP address combinations.
// IP addresses must follow the format A.B.C.D, where A, B, C, and D are numbers between 0 and 255. Zero-prefixed numbers, such as 01 and 065, are not allowed, except for 0 itself.
// For example, given "2542540123", you should return ['254.25.40.123', '254.254.0.123'].

func stringToIP(str string) []string {
	possible := []string{}
	var recurse func(string, int, int)
	l := len(str)

	recurse = func(ip string, start int, oct int) {
		if oct > 3 {
			if start == l {
				possible = append(possible, ip[:len(ip)-1])
			}
			return
		}

		for end := start + 1; end < l+1; end++ {
			num, err := strconv.Atoi(str[start:end])
			if err != nil {
				continue
			}
			if (str[start] == '0' && end-start == 1) || (str[start] != '0' && num > 0 && num < 256) {
				recurse(ip+str[start:end]+".", end, oct+1)
			}
		}
	}

	recurse("", 0, 0)

	return possible
}
