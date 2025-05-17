package main

/*
https://leetcode.com/problems/is-subsequence/
*/
func isSubsequence(subsequence string, text string) bool {
	subIndex := 0
	textIndex := 0
	for textIndex < len(text) && subIndex < len(subsequence) {
		if subsequence[subIndex] == text[textIndex] {
			subIndex++
		}
		textIndex++
	}

	return subIndex == len(subsequence)
}
