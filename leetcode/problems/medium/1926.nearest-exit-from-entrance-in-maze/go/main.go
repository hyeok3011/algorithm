package main

// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/?envType=study-plan-v2&envId=leetcode-75
func nearestExit(maze [][]byte, entrance []int) int {
    maze[entrance[0]][entrance[1]] = '+'
    queue := [][2]int{[2]int{entrance[0], entrance[1]}}
    head := 0
    step := 0
    for head < len(queue) {
        task := len(queue)
        for head < task {
            row, col := queue[head][0], queue[head][1]
            if isExit(row, col, maze) {
                return step
            }

            if row - 1 >= 0 && maze[row - 1][col] == '.' {
                maze[row - 1][col] = '-'
                queue = append(queue, [2]int{row - 1, col})
            }

            if row + 1 < len(maze) && maze[row + 1][col] == '.' {
                maze[row + 1][col] = '-'
                queue = append(queue, [2]int{row + 1, col})
            }

            if col - 1 >= 0 && maze[row][col - 1] == '.' {
                maze[row][col - 1] = '-'
                queue = append(queue, [2]int{row, col - 1})
            }

            if col + 1 < len(maze[row]) && maze[row][col + 1] == '.' {
                maze[row][col + 1] = '-'
                queue = append(queue, [2]int{row, col + 1})
            }
            head++
        }
        step++
    }

    return -1       
}

func isExit(row, col int, maze [][]byte) bool {
	if maze[row][col] == '+' {
		return false
	}

	return (row == 0 || row == len(maze)-1) || (col == 0 || col == len(maze[row])-1)
}
