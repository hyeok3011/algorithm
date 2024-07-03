package main

// https://leetcode.com/problems/valid-palindrome/description/
// Runtime 2 ms Beats 73.67% Memory 4.4 MB Beats 25.36%

func isPalindrome(s string) bool {
	runes := []rune{}

	for _, v := range s {
		if !isPalindromeChar(v) {
			continue
		}

		runes = append(runes, convertToLowerCase(v))
	}

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}

func isPalindromeChar(v rune) bool {
	if 48 <= v && v <= 57 {
		return true
	}
	if 97 <= v && v <= 122 {
		return true
	}

	if 65 <= v && v <= 90 {
		return true
	}
	return false
}

func convertToLowerCase(v rune) rune {
	if 65 <= v && v <= 90 {
		return v + 32
	}

	return v
}

func toLowerCase(c byte) byte {
	if c >= 65 && c <= 90 {
		return c + 32
	}
	return c
}

func isPunctuation(c byte) bool {
	if (c >= 32 && c <= 47) || (c >= 58 && c <= 64) || (c >= 91 && c <= 96) || (c >= 123 && c <= 126) {
		return true
	}
	return false
}

// O(n)으로 메모리 소비도없이 효율적으로 풀었다. 간단한 문제라도 조금 더 생각을 해야겠음
func AnotherSolution(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if isPunctuation(s[i]) {
			i++
			continue
		}

		if isPunctuation(s[j]) {
			j--
			continue
		}

		if toLowerCase(s[i]) == toLowerCase(s[j]) {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}
