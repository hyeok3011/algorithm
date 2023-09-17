package main

func solution(board []string) int {
	oCount := 0
	xCount := 0
	boardArray := [3][3]rune{}
	for i := range board {
		for j, char := range board[i] {
			if char == 'O' {
				oCount++
			} else if char == 'X' {
				xCount++
			}
			boardArray[i][j] = char
		}
	}

	if oCount-xCount > 1 || oCount < xCount {
		return 0
	}

	oBingoCount := getBingoCount('O', boardArray)
	xBingCount := getBingoCount('X', boardArray)
	if oBingoCount > 0 && oBingoCount == xBingCount {
		return 0
	}

	if oBingoCount > 0 && oCount == xCount {
		return 0
	}

	if xBingCount > 0 && oCount > xCount {
		return 0
	}

	return 1
}

func getBingoCount(char rune, board [3][3]rune) int {
	bingoCount := 0

	for i := 0; i < 3; i++ {
		if isBingo(char, board[i]) {
			bingoCount += 1
		}

		if isBingo(char, [3]rune{board[0][i], board[1][i], board[2][i]}) {
			bingoCount += 1
		}
	}

	if isBingo(char, [3]rune{board[0][0], board[1][1], board[2][2]}) {
		bingoCount += 1
	}

	if isBingo(char, [3]rune{board[0][2], board[1][1], board[2][0]}) {
		bingoCount += 1
	}

	return bingoCount
}

func isBingo(char rune, line [3]rune) bool {
	for _, v := range line {
		if char != v {
			return false
		}
	}
	return true
}
