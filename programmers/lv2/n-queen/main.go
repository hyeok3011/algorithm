package main

// https://school.programmers.co.kr/learn/courses/30/lessons/12952
// 뭔가 비슷한 문제를 본듯한 기억이 있다.
// 퀸은 가로, 세로, 대각선 모두 공격이 가능하다.
// 먼저 완전 탐색은 필수일듯하고
// 2차원 배열로도 풀 수 있지만 1차원 배열로하면 훨씬 더 효율적으로 할 수 있다.
// row-col의 절대값을 확인하면 오른쪽 대각선에 중복이 되는지 알 수있다.
// 0 1 2 3
// 1 0 1 2
// 2 1 0 1
// 3 2 1 0
// 순회방법을 개선하면 분명 더 최적화가 가능할거같은데 지금 생각이 나지 않는다.
const (
	UNPLACED = 9999
)

func solution(n int) int {
	chessBoard := make([]int, n)
	for i := range chessBoard {
		chessBoard[i] = UNPLACED
	}
	count := 0

	count = backTracking(chessBoard, 0, n, count)
	return count
}

func backTracking(chessBoard []int, row, n, count int) int {
	if row == n {
		count += 1
		return count
	}

	for i := 0; i < n; i++ {
		if isValidPlacement(chessBoard, row, i) {
			chessBoard[row] = i
			count = backTracking(chessBoard, row+1, n, count)
			chessBoard[row] = UNPLACED
		}
	}
	return count
}

func isValidPlacement(chessBoard []int, row, col int) bool {
	for i, v := range chessBoard {
		if v == col {
			return false
		}

		if abs(i-row) == abs(v-col) {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}
