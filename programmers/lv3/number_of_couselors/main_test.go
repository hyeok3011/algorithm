package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	input := [][]int{{10, 60, 1}, {15, 100, 3}, {20, 30, 1}, {30, 50, 3}, {50, 40, 1}, {60, 30, 2}, {65, 30, 1}, {70, 100, 2}}
	assert.Equal(t, solution(3, 5, input), 25)
	input2 := [][]int{{5, 55, 2}, {10, 90, 2}, {20, 40, 2}, {50, 45, 2}, {100, 50, 2}}
	assert.Equal(t, solution(2, 3, input2), 90)

	input3 := [][]int{{5, 55, 1}, {10, 90, 2}, {20, 40, 3}}
	assert.Equal(t, solution(3, 3, input3), 0)
}

// 참가자 번호	시각	상담 시간	상담 유형
// 1번 참가자	10분	60분	1번 유형
// 2번 참가자	15분	100분	3번 유형
// 3번 참가자	20분	30분	1번 유형
// 4번 참가자	30분	50분	3번 유형
// 5번 참가자	50분	40분	1번 유형
// 6번 참가자	60분	30분	2번 유형
// 7번 참가자	65분	30분	1번 유형
// 8번 참가자	70분	100분	2번 유형

// 참가자 번호	시각	상담 시간	상담 유형
// 1번 참가자	5분	   55분	     2번 유형
// 2번 참가자	10분	90분	2번 유형
// 3번 참가자	20분	40분	2번 유형
// 4번 참가자	50분	45분	2번 유형
// 5번 참가자	100분	50분
