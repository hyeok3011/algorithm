package main

// https://leetcode.com/problems/plus-one/description/
// Runtime 0 ms Beats 100% Memory 2.1 MB Beats 54.11%

func plusOne(digits []int) []int {
	currentIndex := len(digits) - 1
	digits[currentIndex] += 1

	for digits[currentIndex] > 9 {
		digits[currentIndex] = 0

		currentIndex -= 1
		if currentIndex < 0 {
			digits = append([]int{0}, digits...)
			currentIndex = 0
		}

		digits[currentIndex] += 1

	}
	return digits
}
