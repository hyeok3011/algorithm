package main

import "sort"

// https://school.programmers.co.kr/learn/courses/30/lessons/154540

func solution(maps []string) []int {

	answer := []int{}
	board := make([][]int, len(maps))
	for i := range maps {
		board[i] = make([]int, len(maps[i]))
		for j, v := range maps[i] {
			if v == 'X' {
				board[i][j] = 0
			} else {
				board[i][j] = int(maps[i][j] - '0')
			}
		}
	}
	maxX := len(maps) - 1
	maxY := len(maps[0]) - 1
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				continue
			}
			queue := [][]int{{i, j}}
			total := 0
			for len(queue) > 0 {
				queueLen := len(queue)
				for k := 0; k < queueLen; k++ {
					x, y := queue[k][0], queue[k][1]
					if board[x][y] == 0 {
						continue
					}

					total += board[x][y]
					board[x][y] = 0
					if x > 0 && board[x-1][y] != 0 {
						queue = append(queue, []int{x - 1, y})
					}

					if x < maxX && board[x+1][y] != 0 {
						queue = append(queue, []int{x + 1, y})
					}

					if y > 0 && board[x][y-1] != 0 {
						queue = append(queue, []int{x, y - 1})
					}

					if y < maxY && board[x][y+1] != 0 {
						queue = append(queue, []int{x, y + 1})
					}
				}

				queue = queue[queueLen:]
			}
			answer = append(answer, total)
		}
	}

	if len(answer) == 0 {
		answer = append(answer, -1)
	}

	sort.Ints(answer)
	return answer
}
