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
