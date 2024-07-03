package main

import "log"

const (
	INT_MAXIMUM           = 2147483647
	FIBONACCI_MINIMUM_LEN = 3
)

// https://leetcode.com/problems/split-array-into-fibonacci-sequence/description/
// Runtime 0 ms Beats 100% Sorry, there are not enough accepted submissions to show data Memory 2 MB Beats 100%
// @@@@@@@@ 또 볼것 @@@@@@@@
func splitIntoFibonacci(num string) []int {
	result := make([]int, 0)
	backtracking(num, 0, &result)
	return result
}

func backtracking(num string, numStartIndex int, result *[]int) bool {
	// isCompleted
	if len(*result) >= FIBONACCI_MINIMUM_LEN && numStartIndex == len(num) {
		return true
	}

	for currentIndex := numStartIndex; currentIndex < len(num); currentIndex++ {
		// isStartZero
		if num[numStartIndex] == '0' && currentIndex > numStartIndex {
			break
		}

		number := convertStringToInt(num[numStartIndex : currentIndex+1])
		if number > INT_MAXIMUM {
			break
		}

		resultLen := len(*result)
		if resultLen >= 2 && number > (*result)[resultLen-1]+(*result)[resultLen-2] {
			break
		}

		if resultLen <= 1 || number == (*result)[resultLen-1]+(*result)[resultLen-2] {
			*result = append(*result, number)
			if backtracking(num, currentIndex+1, result) {
				return true
			}
			*result = (*result)[:resultLen]
		}
		log.Println("temp")
	}

	return false
}

func convertStringToInt(strNumber string) int {
	num := 0
	for _, v := range strNumber {
		num = num*10 + int(v-'0')
	}
	return num
}
