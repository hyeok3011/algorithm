package main
// https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
    var rowMask, colMask, subBoxMask [9][9]bool
    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            if board[row][col] == '.' {
                continue
            }
            number := board[row][col] - '1'
            boxRow := (col / 3) + (row / 3 * 3)
            if rowMask[row][number] || colMask[col][number] || subBoxMask[boxRow][number] {
                return false
            }
            rowMask[row][number] = true
            colMask[col][number] = true
            subBoxMask[boxRow][number] = true
        }
    }
    return true
}

// 처음에 간단히 풀었던 코드, 하나의 루프로 통합하여 다시 풀어봄
func AnotherSolution(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if hasDuplicateInBox(i, board) || hasDuplicateInRow(i, board) || hasDuplicateInCol(i, board) {
			return false
		}
	}
	return true
}

func hasDuplicateInBox(boxIndex int, board [][]byte) bool {
	startRow := (boxIndex / 3) * 3
	startCol := (boxIndex % 3) * 3
	seen := make([]bool, 9)
	for row := startRow; row < startRow+3; row++ {
		for col := startCol; col < startCol+3; col++ {
			if board[row][col] == '.' {
				continue
			}
			num := board[row][col] - '1'
			if seen[num] {
				return true
			}
			seen[num] = true
		}
	}
	return false
}

func hasDuplicateInRow(row int, board [][]byte) bool {
	seen := make([]bool, 9)
	for col := 0; col < 9; col++ {
		if board[row][col] == '.' {
			continue
		}
		num := board[row][col] - '1'
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func hasDuplicateInCol(col int, board [][]byte) bool {
	seen := make([]bool, 9)
	for row := 0; row < 9; row++ {
		if board[row][col] == '.' {
			continue
		}
		num := board[row][col] - '1'
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}
