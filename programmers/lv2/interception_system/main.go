package main

// https://school.programmers.co.kr/learn/courses/30/lessons/181188
import (
	"sort"
)

func solution(targets [][]int) int {
	if len(targets) == 1 {
		return 1
	}

	sort.Slice(targets, func(i, j int) bool {
		return targets[i][0] < targets[j][0]
	})

	interceptionCount := 1
	currentIndex := 1
	minX := targets[0][0]
	maxX := targets[0][1]
	for currentIndex < len(targets) {
		if minX <= targets[currentIndex][0] && targets[currentIndex][0] < maxX {
			if targets[currentIndex][1] < maxX {
				maxX = targets[currentIndex][1]
			}
		} else {
			interceptionCount += 1
			minX = targets[currentIndex][0]
			maxX = targets[currentIndex][1]
		}
		currentIndex += 1
	}

	return interceptionCount
}

// 반대 정렬 후 시작점을 보면 더 쉬워질수도 있음
func solution2(targets [][]int) int {
	if len(targets) == 1 {
		return 1
	}

	sort.Slice(targets, func(i, j int) bool {
		return targets[i][0] > targets[j][0]
	})

	interceptionCount := 1
	currentIndex := 1
	startPoint := targets[0][0]
	// log.Println(targets)
	for currentIndex < len(targets) {
		if startPoint >= targets[currentIndex][1] {
			startPoint = targets[currentIndex][0]
			interceptionCount += 1
		}
		currentIndex += 1
	}

	return interceptionCount
}
