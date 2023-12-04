package main

/*
 */
const (
	MAX_PRICE = 400000
)

func Solution(A []int) int {
	// Implement your solution here
	boughtPrice := MAX_PRICE
	maxProfit := 0
	for _, v := range A {
		if boughtPrice > v {
			boughtPrice = v
		} else {
			maxProfit = max(maxProfit, v-boughtPrice)
		}
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
