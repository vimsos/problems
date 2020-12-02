package dcp

import "fmt"

// The "look and say" sequence is defined as follows: beginning with the term 1, each subsequent term visually describes the digits appearing in the previous term. The first few terms are as follows:
// {1, 11, 21, 1211, 111221}
// As an example, the fourth term is 1211, since the third term consists of one 2 and one 1.
// Given an integer N, print the Nth term of this sequence.

func lookAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	cur, prev := "", "1"
	for term := 1; term < n; term++ {
		qtd := 1
		cur = ""
		for i := 0; i < len(prev)-1; i++ {
			if prev[i] == prev[i+1] {
				qtd++
			} else {
				cur = cur + fmt.Sprint(qtd) + string(prev[i])
				qtd = 1
			}
		}
		cur = cur + fmt.Sprint(qtd) + string(prev[len(prev)-1])

		prev = cur
	}

	return cur
}
