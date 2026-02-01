package main

// https://leetcode.com/problems/palindrome-partitioning/description/
// isPalindrome에 memoization을 적용하면 작지만 조금 더 최적화는 가능한 상황
func partition(s string) [][]string {
	result := [][]string{}
	recursion(&result, 0, []string{}, s)
	return result
}

func recursion(result *[][]string, start int, current []string, s string) {
	if start >= len(s) {
		temp := make([]string, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(start, i, s) {
			current = append(current, string(s[start:i+1]))
			recursion(result, i+1, current, s)
			current = current[:len(current)-1]
		}
	}
}

func isPalindrome(left, right int, s string) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
