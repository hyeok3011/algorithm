package main

// https://school.programmers.co.kr/learn/courses/30/lessons/154538
func solution(x int, y int, n int) int {
	if x == y {
		return 0
	}

	step := 0
	queue := []int{x * 2, x * 3, x + n}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		currentDepthLen := len(queue)
		step += 1
		for i := 0; i < currentDepthLen; i++ {
			currentVal := queue[i]
			if currentVal > y || visited[currentVal] {
				continue
			}

			if currentVal == y {
				return step
			}
			visited[currentVal] = true

			queue = append(queue, currentVal*2)
			queue = append(queue, currentVal*3)
			queue = append(queue, currentVal+n)
		}

		queue = queue[currentDepthLen:]
	}
	return -1
}
