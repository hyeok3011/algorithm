package main

// https://leetcode.com/problems/01-matrix/submissions/1847640763/
// @@@@@@@
// 처음 접근을 bfs + 메모이제이션 방식으로 풀어보려고 했으나 실패함.
// 풀고 다른 사람의 풀이를 보니 충분히 가능한것같아. 다시 풀어봐야함.
func updateMatrix(mat [][]int) [][]int {
	for row := range mat {
		for col := range mat[row] {
			if mat[row][col] == 1 {
				mat[row][col] = 9999999
			}
		}
	}

	for row := range mat {
		for col := range mat[row] {
			if row+1 < len(mat) {
				mat[row+1][col] = min(mat[row][col]+1, mat[row+1][col])
			}

			if col+1 < len(mat[row]) {
				mat[row][col+1] = min(mat[row][col]+1, mat[row][col+1])
			}
		}
	}

	for row := len(mat) - 1; row >= 0; row-- {
		for col := len(mat[row]) - 1; col >= 0; col-- {
			if row-1 >= 0 {
				mat[row-1][col] = min(mat[row][col]+1, mat[row-1][col])
			}

			if col-1 >= 0 {
				mat[row][col-1] = min(mat[row][col]+1, mat[row][col-1])
			}
		}
	}

	return mat
}
