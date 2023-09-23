package main

func solution(maps []string) int {
	var currentX, currentY int
	visited := make([][]bool, len(maps))
	for i := range maps {
		for j := range maps[i] {
			if maps[i][j] == 'S' {
				currentX = i
				currentY = j
			}
		}
		visited[i] = make([]bool, len(maps[i]))
	}
	leverX, leverY, leverDuration := bfs(currentX, currentY, 'L', maps, visited)
	for i := range visited {
		for j := range visited[i] {
			visited[i][j] = false
		}
	}
	if leverDuration == -1 {
		return -1
	}

	_, _, exitDuration := bfs(leverX, leverY, 'E', maps, visited)
	if exitDuration == -1 {
		return -1
	}

	return leverDuration + exitDuration
}

func bfs(currentX, currentY int, target rune, board []string, visited [][]bool) (int, int, int) {
	minX := 0
	maxX := len(board) - 1
	minY := 0
	maxY := len(board[0]) - 1

	queue := [][]int{
		{currentX, currentY},
	}

	duration := -1
	notFound := true
	for len(queue) > 0 && notFound {
		duration += 1
		currentDurationLen := len(queue)
		for i := 0; i < currentDurationLen; i++ {
			currentX, currentY = queue[i][0], queue[i][1]
			if visited[currentX][currentY] {
				continue
			}

			visited[currentX][currentY] = true
			if board[currentX][currentY] == byte(target) {
				notFound = false
				break
			}

			if currentY > minY && board[currentX][currentY-1] != 'X' {
				queue = append(queue, []int{currentX, currentY - 1})
			}

			if currentY < maxY && board[currentX][currentY+1] != 'X' {
				queue = append(queue, []int{currentX, currentY + 1})
			}

			if currentX > minX && board[currentX-1][currentY] != 'X' {
				queue = append(queue, []int{currentX - 1, currentY})
			}

			if currentX < maxX && board[currentX+1][currentY] != 'X' {
				queue = append(queue, []int{currentX + 1, currentY})
			}

		}
		queue = queue[currentDurationLen:]
	}

	if notFound {
		duration = -1
	}

	return currentX, currentY, duration
}
