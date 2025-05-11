package main

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	reversed := []rune{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		j := i
		for j >= 0 && s[j] != ' ' {
			j--
		}
		reversed = append(reversed, []rune(s[j+1:i+1])...)
		reversed = append(reversed, ' ')
		i = j
	}
	return string(reversed[:len(reversed)-1])
}

func reverseWords2(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}
