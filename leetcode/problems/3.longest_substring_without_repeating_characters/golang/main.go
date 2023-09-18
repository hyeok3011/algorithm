package main

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	substring := make([]byte, 128)
	maxSubStringLength := 0
	for end < len(s) {
		if substring[s[end]] == 0 {
			substring[s[end]] = 1
			end += 1
			maxSubStringLength = max(maxSubStringLength, end-start)
		} else {
			substring[s[start]] = 0
			start += 1
		}
	}

	return maxSubStringLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
