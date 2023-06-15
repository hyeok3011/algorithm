package main

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// Runtime 2 ms Beats 60.67% Memory 2 MB Beats 66.52%
const (
	DID_NOT_OCCUR = -1
)

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if string(haystack[i:len(needle)+i]) == needle {
			return i
		}
	}

	return DID_NOT_OCCUR
}
