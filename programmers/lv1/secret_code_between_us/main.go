package main

import "strings"

// https://school.programmers.co.kr/learn/courses/30/lessons/155652

func solution(s string, skip string, index int) string {
	var builder strings.Builder

	skipHashMap := make(map[rune]bool)

	for _, v := range skip {
		skipHashMap[v] = true
	}

	for _, v := range s {
		i := 0
		for i < index {
			v += 1
			if v > 'z' {
				v = 'a'
			}
			if _, exist := skipHashMap[v]; !exist {
				i++
			}
		}
		builder.WriteRune(v)
	}
	return builder.String()
}
