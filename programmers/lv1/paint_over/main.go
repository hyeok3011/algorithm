package main

func solution(n int, m int, section []int) int {

	drawCount := 0
	roller := 0

	for _, v := range section {
		if roller >= v {
			continue
		}
		roller = m + v - 1
		drawCount += 1
	}
	return drawCount
}
