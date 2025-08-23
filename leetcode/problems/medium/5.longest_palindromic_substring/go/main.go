package main

/*
https://leetcode.com/problems/longest-palindromic-substring/solutions/2752538/golang-explicado-video/
*/
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	var left, right int
	answer := s[:1]
	center := 0
	for center < len(s) {
		left, right = center, center

		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}

		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}

		if len(answer) < right-left+1 {
			answer = s[left : right+1]
		}
		center = (left+right)/2 + 1
	}

	return answer
}
