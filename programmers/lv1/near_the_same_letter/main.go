package main

// https://school.programmers.co.kr/learn/courses/30/lessons/142086
// 간단하다 글자들의 인덱스를 맵에 저장해놓고 계산하면 된다.
// array와 map의 사이즈를 정해놓ㅇ를 수 있으면 항상 정해놓는 습관을 들이자
func solution(s string) []int {
	answer := make([]int, len(s))
	letterIndex := make(map[rune]int, 26)
	for index, v := range s {
		if preIndex, exist := letterIndex[v]; exist {
			answer[index] = index - preIndex
		} else {
			answer[index] = -1
		}
		letterIndex[v] = index
	}

	return answer
}
