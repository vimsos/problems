package leetcode

import "math"

// Say you have an array for which the ith element is the price of a given stock on day i.
// If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
// Note that you cannot sell a stock before you buy one.

func maxProfit(prices []int) int {
	profit, minPrice := 0, math.MaxInt32

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if p := prices[i] - minPrice; p > profit {
			profit = p
		}
	}

	return profit
}
