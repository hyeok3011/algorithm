package main

/*
https://leetcode.com/problems/longest-palindromic-substring/solutions/2752538/golang-explicado-video/


*/

func longestPalindrome(s string) string {
	var left, right, palindromeLeft, palindromeRight int
	for right < len(s) {
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}

		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}

		if right-left > palindromeRight-palindromeLeft {
			palindromeLeft, palindromeRight = left, right
		}

		left = (left+right)/2 + 1
		right = left
	}
	return s[palindromeLeft : palindromeRight+1]
}
