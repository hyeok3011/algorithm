package main

/*
https://leetcode.com/problems/is-subsequence/
*/

func isSubsequence(s string, t string) bool {
	subIndex := 0
	for _, v := range t {
		if subIndex < len(s) && v == rune(s[subIndex]) {
			subIndex++
		}
	}

	return subIndex <= len(s)
}
