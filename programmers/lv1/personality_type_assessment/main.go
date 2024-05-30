package main

import "strings"

// https://school.programmers.co.kr/learn/courses/30/lessons/118666
// 문제만 잘 이해된다면 풀기 쉬운 문제
// RT TR 문자 순서에 따라서 점수개념만
func solution(survey []string, choices []int) string {
	var personality strings.Builder
	personalityScoreMap := map[rune]int{'R': 0, 'T': 0, 'C': 0, 'F': 0, 'J': 0, 'M': 0, 'A': 0, 'N': 0}

	for i, v := range survey {
		if choices[i] >= 4 {
			personalityScoreMap[rune(v[1])] += choices[i] - 4
		} else {
			personalityScoreMap[rune(v[0])] += 4 - choices[i]
		}
	}

	for _, pair := range [][2]rune{{'R', 'T'}, {'C', 'F'}, {'J', 'M'}, {'A', 'N'}} {
		if personalityScoreMap[pair[0]] >= personalityScoreMap[pair[1]] {
			personality.WriteRune(pair[0])
		} else {
			personality.WriteRune(pair[1])
		}
	}

	return personality.String()
}
