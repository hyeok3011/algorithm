package main

import "sort"

// https://school.programmers.co.kr/learn/courses/30/lessons/138477
func solution(k int, scores []int) []int {
	answer := []int{}

	fameQueue := []int{}
	for _, score := range scores {
		fameQueue = append(fameQueue, score)
		sort.Ints(fameQueue)
		if k < len(fameQueue) {
			fameQueue = fameQueue[1:]
		}

		answer = append(answer, fameQueue[0])
	}
	return answer
}
