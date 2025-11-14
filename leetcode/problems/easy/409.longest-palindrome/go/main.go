package main

// https://leetcode.com/problems/longest-palindrome/description/
// palindrome은 홀수 종류를 1개, 그 외에는 모두 짝수로만 존재해야 한다.
func longestPalindrome(s string) int {
	letterCount := make(map[rune]int)
	for _, v := range s {
		if _, exist := letterCount[v]; !exist {
			letterCount[v] = 0
		}
		letterCount[v]++
	}

	totalCount := 0
	hasOdd := false
	for _, value := range letterCount {
		if value%2 == 1 {
			totalCount += value - 1
			hasOdd = true
		} else {
			totalCount += value
		}
	}

	if hasOdd {
		totalCount++
	}
	return totalCount
}
