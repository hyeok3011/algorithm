package main

const (
	START                = 'S'
	ACCESSIBLE_PASSAGE   = 'O'
	INACCESSIBLE_PASSAGE = 'X'
)

// https://school.programmers.co.kr/learn/courses/30/lessons/172928
// 어렵지는 않지만 풀다가 짜증났음...
func solution(park []string, routes []string) []int {
	currentPostion := make([]int, 2)
	for i, v := range park {
		for j, char := range v {
			if char == START {
				currentPostion[0] = i
				currentPostion[1] = j
			}
		}
	}

	for _, route := range routes {
		positionIndex, moveCount := routeToPositionDelta(route)
		if isAccessiblePASSAGE(currentPostion, positionIndex, moveCount, park) {
			currentPostion[positionIndex] += moveCount
		}
	}

	return currentPostion
}

func routeToPositionDelta(route string) (int, int) {
	direction := rune(route[0])
	moveCount := int(route[2] - '0')

	switch direction {
	case 'N':
		return 0, -moveCount
	case 'S':
		return 0, moveCount
	case 'W':
		return 1, -moveCount
	case 'E':
		return 1, moveCount
	}
	return 0, 0
}

func isAccessiblePASSAGE(currentPostion []int, positionIndex int, moveCount int, park []string) bool {
	nextPosition := make([]int, 2)
	nextPosition[0] = currentPostion[0]
	nextPosition[1] = currentPostion[1]

	increment := 1
	if moveCount < 0 {
		increment = -1
	}

	for i := 0; i != moveCount; i += increment {
		nextPosition[positionIndex] += increment

		if nextPosition[0] < 0 || nextPosition[0] >= len(park) || nextPosition[1] < 0 || nextPosition[1] >= len(park[0]) {
			return false
		}

		if park[nextPosition[0]][nextPosition[1]] == INACCESSIBLE_PASSAGE {
			return false
		}
	}
	return true
}
