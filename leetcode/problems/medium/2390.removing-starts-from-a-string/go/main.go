package main

// https://leetcode.com/problems/removing-stars-from-a-string/?envType=study-plan-v2&envId=leetcode-75
func removeStars(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

// fast - slow 방식 in-place방식
func removeStars(s string) string {
    fast, slow := 0, 0
    result := make([]byte, len(s))
    for fast < len(s) {
        if s[fast] == '*' {
            slow--
        } else {
            result[slow] = s[fast]
            slow++
        }
        fast++
    }
    return string(result[:slow])
}
