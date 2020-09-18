package leetcode

// Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers. The solution set must not contain duplicate combinations.

func combinationSum3(k int, n int) [][]int {
	var dfs func(int, []int, *[][]int)
	dfs = func(min int, set []int, combs *[][]int) {
		if len(set) == k {
			sum := 0
			for i := 0; i < k; i++ {
				sum += set[i]
			}
			if sum == n {
				*combs = append(*combs, set)
			}
			return
		}
		for i := min; i < 10; i++ {
			nextSet := []int{}
			nextSet = append(nextSet, set...)
			nextSet = append(nextSet, i)

			dfs(i+1, nextSet, combs)
		}
	}

	var combinations [][]int
	dfs(1, []int{}, &combinations)
	return combinations
}

// 	// brian kernighan’s algorithm
// 	countSetBits := func(num int) int {
// 		count := 0
// 		for num > 0 {
// 			num &= num - 1
// 			count++
// 		}
// 		return count
// 	}

// 	intToSet := func(num int) (int, []int) {
// 		sum, pos, index := 0, 1, 0
// 		set := make([]int, k)
// 		for num > 0 && pos < 10 {
// 			isSet := num & 1
// 			if isSet == 1 {
// 				set[index] = pos
// 				sum += pos
// 				index++
// 			}
// 			pos++
// 			num = num >> 1
// 		}
// 		return sum, set
// 	}

// 	total := 9
// 	for i := 1; i < k; i++ {
// 		total *= 9 - i
// 	}
// 	combinations := [][]int{}
// 	for i := 0; i < total; i++ {
// 		if countSetBits(i) == k {
// 			sum, set := intToSet(i)
// 			if sum == n {
// 				combinations = append(combinations, set)
// 			}
// 		}
// 	}

// 	return combinations
// }

// func combinationSum3(k int, n int) [][]int {
// 	var combinations [][]int
// 	set := make([]int, k)
// 	checkSet := func(s []int) {
// 		var sum int
// 		for i := 0; i < k; i++ {
// 			sum += s[i]
// 		}
// 		if sum == n {
// 			combinations = append(combinations, s)
// 		}
// 	}

// 	// initialize set with the first possible combination
// 	for i := 0; i < k; i++ {
// 		set[i] = i + 1
// 	}
// 	checkSet(set)

// 	for
// 	for i := k - 1; i > 0; i-- {
// 		fmt.Println(set, " ")
// 		set[i] = (set[i-1] + 1) % 9
// 	}

// 	return combinations
// }
