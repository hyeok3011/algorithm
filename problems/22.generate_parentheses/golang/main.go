package main

// https://leetcode.com/problems/generate-parentheses/description/
// Runtime 0 ms Beats 100% Memory 2.7 MB Beats 53.92%
// dfs
func generateParenthesis(n int) []string {
	result := []string{}
	combination("(", 1, 0, n, &result)
	return result
}

func combination(current string, openCount, closeCount, n int, result *[]string) {
	if openCount == n && closeCount == n {
		*result = append(*result, current)
		return
	}

	if openCount < n {
		combination(current+"(", openCount+1, closeCount, n, result)
	}

	if closeCount < n && openCount > closeCount {
		combination(current+")", openCount, closeCount+1, n, result)
	}
}
