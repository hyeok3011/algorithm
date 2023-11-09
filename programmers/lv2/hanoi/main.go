package main

// https://school.programmers.co.kr/learn/courses/30/lessons/12946
// 너무 유명한 하노이 문제다.

func solution(n int) [][]int {
	return hanoi(n, 1, 3, 2, [][]int{})
}

func hanoi(disk, start, target, aux int, answer [][]int) [][]int {
	if disk == 1 {
		return append(answer, []int{start, target})
	} else {
		answer = hanoi(disk-1, start, aux, target, answer)
		answer = append(answer, []int{start, target})
		answer = hanoi(disk-1, aux, target, start, answer)
	}
	return answer
}
