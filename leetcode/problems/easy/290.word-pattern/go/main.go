package main

// https://leetcode.com/problems/word-pattern/description/
func wordPattern(pattern string, s string) bool {
	charMap := make(map[byte]string)
	wordMap := make(map[string]byte)
	left := 0
	right := 0
	for i := range pattern {
		if right >= len(s) {
			return false
		}
		for right < len(s) && s[right] != ' ' {
			right++
		}

		char := pattern[i]
		word := string(s[left:right])

		if mapped, ok := charMap[char]; ok && word != mapped {
			return false
		}

		if mapped, ok := wordMap[word]; ok && char != mapped {
			return false
		}

		charMap[char] = word
		wordMap[word] = char
		left = right + 1
		right = left
	}
	return right > len(s)
}
