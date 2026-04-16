package main

// https://leetcode.com/problems/perfect-squares/
// 최소를 구해야하기때문에 큰 숫자부터 접근하는 greedy하게 접근할수가 없다. 예외 상황이 존재하기 떄문이다.
// @@@@@@ DP로 풀다가 이상하게 헷갈렸던 문제. 다시 풀어보면 좋을듯.
func numSquares(n int) int {
	squares := []int{}
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	visited := make([]bool, n+1)
	visited[n] = true

	queue := []int{n}
	level := 0

	for len(queue) > 0 {
		level++
		nextQueue := []int{}

		for _, curr := range queue {
			for _, num := range squares {
				next := curr - num

				if next < 0 {
					break
				}
				if next == 0 {
					return level
				}
				if visited[next] {
					continue
				}
				visited[next] = true
				nextQueue = append(nextQueue, next)
			}
		}

		queue = nextQueue
	}

	return level
}
