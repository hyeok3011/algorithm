package main

/*
https://leetcode.com/problems/word-search/description/
음... 단순한 백트래킹 문제로 보인다.
*/

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))

	for i := range board {
		visited[i] = make([]bool, len(board[i]))
	}

	for row := range board {
		for col := range board[row] {
			if dfs(board, visited, row, col, 0, word) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, visited [][]bool, row, col, wordIndex int, word string) bool {
	if wordIndex == len(word)-1 {
		return board[row][col] == word[wordIndex]
	}

	if board[row][col] != word[wordIndex] {
		return false
	}

	visited[row][col] = true
	if row-1 >= 0 && !visited[row-1][col] && dfs(board, visited, row-1, col, wordIndex+1, word) {
		return true
	}

	if row+1 < len(board) && !visited[row+1][col] && dfs(board, visited, row+1, col, wordIndex+1, word) {
		return true
	}

	if col-1 >= 0 && !visited[row][col-1] && dfs(board, visited, row, col-1, wordIndex+1, word) {
		return true
	}

	if col+1 < len(board[0]) && !visited[row][col+1] && dfs(board, visited, row, col+1, wordIndex+1, word) {
		return true
	}

	visited[row][col] = false
	return false
}
