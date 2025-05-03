package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var stringBuilder strings.Builder

	minLength := min(len(word1), len(word2))

	for i := 0; i < minLength; i++ {
		stringBuilder.WriteByte(word1[i])
		stringBuilder.WriteByte(word2[i])
	}

	if len(word1) > minLength {
		stringBuilder.WriteString(word1[minLength:])
	} else {
		stringBuilder.WriteString(word2[minLength:])
	}

	return stringBuilder.String()
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
