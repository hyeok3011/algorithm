package main

// https://leetcode.com/problems/surrounded-regions/description/

func solve(board [][]byte) {
	var recursion func(row, col int)
	recursion = func(row, col int) {
		if row == len(board) || col == len(board[0]) || row < 0 || col < 0 {
			return
		}

		if board[row][col] == 'X' || board[row][col] == 'T' {
			return
		}

		board[row][col] = 'T'
		recursion(row+1, col)
		recursion(row-1, col)
		recursion(row, col+1)
		recursion(row, col-1)
	}

	for col := range board[0] {
		recursion(0, col)
		recursion(len(board)-1, col)
	}

	for row := 0; row < len(board); row++ {
		recursion(row, 0)
		recursion(row, len(board[0])-1)
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			} else if board[row][col] == 'T' {
				board[row][col] = 'O'
			}
		}
	}
}
