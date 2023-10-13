package main

import (
	"sort"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/135808
// 한 박스에 들어가는 과일중 가장 낮은 가격을 기준으로 과일 개수를 곱하면 된다.
// k는 왜 주는지 사실 잘 모르겠다.
func solution(k int, boxSize int, score []int) int {
	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	totalBenefit := 0
	boxCount := len(score) / boxSize
	for i := 0; i < boxCount; i++ {
		totalBenefit += score[(i+1)*boxSize-1] * boxSize
	}

	return totalBenefit
}
