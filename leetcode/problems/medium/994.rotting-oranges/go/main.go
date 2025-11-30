package main


// https://leetcode.com/problems/rotting-oranges/description/?envType=problem-list-v2&envId=rabvlt31
func orangesRotting(grid [][]int) int {
	rottenOrange := [][]int{}
	freshCount := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 2 {
				rottenOrange = append(rottenOrange, []int{row, col})
			} else if grid[row][col] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	minute := 0

	for len(rottenOrange) > 0 && freshCount > 0 {
		currentTaskLen := len(rottenOrange)
		for i := 0; i < currentTaskLen; i++ {
			position := rottenOrange[i]
			row, col := position[0], position[1]

			if row+1 < len(grid) && grid[row+1][col] == 1 {
				rottenOrange = append(rottenOrange, []int{row + 1, col})
				grid[row+1][col] = 2
				freshCount--
			}

			if row-1 >= 0 && grid[row-1][col] == 1 {
				rottenOrange = append(rottenOrange, []int{row - 1, col})
				grid[row-1][col] = 2
				freshCount--
			}

			if col+1 < len(grid[row]) && grid[row][col+1] == 1 {
				rottenOrange = append(rottenOrange, []int{row, col + 1})
				grid[row][col+1] = 2
				freshCount--
			}

			if col-1 >= 0 && grid[row][col-1] == 1 {
				rottenOrange = append(rottenOrange, []int{row, col - 1})
				grid[row][col-1] = 2
				freshCount--
			}
		}

		rottenOrange = rottenOrange[currentTaskLen:]

        minute++
	}

	if freshCount > 0 {
		return -1
	}

	return minute
}
