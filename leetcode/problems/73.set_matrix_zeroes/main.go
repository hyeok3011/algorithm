package main

/*
https://leetcode.com/problems/set-matrix-zeroes/
굉장히 간단한 문제이다.
matrix[x][y]가 0이라면 좌표기준으로 십자가 모양으로 모두 0으로 만들어주면 된다.
단순하게 0인경우 x와 y for문을 중복으로 돌지 않도록 index를 따로 맵에 저장하여
for loop를 도는식으로 구현했다.
*/

func setZeroes(matrix [][]int) {
	rowIndexes := make(map[int]struct{})
	colIndexes := make(map[int]struct{})
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				rowIndexes[row] = struct{}{}
				colIndexes[col] = struct{}{}
			}
		}
	}

	for row := range rowIndexes {
		for col := range matrix[row] {
			matrix[row][col] = 0
		}
	}

	for col := range colIndexes {
		for row := 0; row < len(matrix); row++ {
			matrix[row][col] = 0
		}
	}
}
