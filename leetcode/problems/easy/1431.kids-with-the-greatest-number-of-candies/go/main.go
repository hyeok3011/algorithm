package main

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description/?envType=study-plan-v2&envId=leetcode-75
func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	maxCandies := max(candies)
	threshold := maxCandies - extraCandies

	for i, candyCount := range candies {
		result[i] = candyCount >= threshold
	}
	return result
}

func max(elements []int) int {
	max := elements[0]
	for _, elem := range elements {
		if max < elem {
			max = elem
		}
	}
	return max
}
