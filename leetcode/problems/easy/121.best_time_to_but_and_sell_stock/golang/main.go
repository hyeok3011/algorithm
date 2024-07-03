package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
// Runtime 101 ms Beats 99.46% Memory 7.9 MB Beats 94.41%
func maxProfit(prices []int) int {
	minPrice := prices[0]
	bestProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			if bestProfit < profit {
				bestProfit = profit
			}
		}
	}
	return bestProfit
}
