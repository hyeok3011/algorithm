package main

// https://school.programmers.co.kr/learn/courses/30/lessons/77485
// 어.... 분명 처음 문제봤을때 최악의경우 100 * 100 행렬에서 10000번을 돌려야해서
// 더 좋은 방법이 있나 아무리생각해도 정답이 떠오르지 않아 무작정 풀어보니 그냥 통과가 된다...
// 시계방향으로 밀때 조금 헷갈렸지만 딱히 어려운문제같지는 않다.
// 좋은 방법으로 푼 사람도 없어 보인다.
func solution(rows int, columns int, queries [][]int) []int {
	matrix := make([][]int, rows)
	cnt := 1
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			matrix[i][j] = cnt
			cnt++
		}
	}
	answer := make([]int, len(queries))
	for i := range queries {
		answer[i] = rotateAndFindMin(matrix, queries[i][0]-1, queries[i][1]-1, queries[i][2]-1, queries[i][3]-1)
	}
	return answer
}

func rotateAndFindMin(matrix [][]int, x1, y1, x2, y2 int) int {
	tmp := matrix[x1][y1]
	minVal := tmp

	// left to top
	for i := x1; i < x2; i++ {
		matrix[i][y1] = matrix[i+1][y1]
		minVal = min(minVal, matrix[i][y1])
	}

	// bottom to left
	for j := y1; j < y2; j++ {
		matrix[x2][j] = matrix[x2][j+1]
		minVal = min(minVal, matrix[x2][j])
	}

	// right to bottom
	for i := x2; i > x1; i-- {
		matrix[i][y2] = matrix[i-1][y2]
		minVal = min(minVal, matrix[i][y2])
	}

	// top to right
	for j := y2; j > y1+1; j-- {
		matrix[x1][j] = matrix[x1][j-1]
		minVal = min(minVal, matrix[x1][j])
	}

	matrix[x1][y1+1] = tmp

	return minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
