package main

// https://school.programmers.co.kr/learn/courses/30/lessons/389481?language=go
// 진짜 쉬울줄 알고 풀었는데 생각보다 오래걸린 문제....
// 특히 숫자에서 알파벳으로 다시 전환할때 시간이 조금 걸렸음..
import (
	"math"
	"sort"
)

func solution(n int64, bans []string) string {
	bansMagicNumber := make([]int, len(bans))
	for i := 0; i < len(bans); i++ {
		bansMagicNumber[i] = calculateMagicNumber(bans[i])
	}
	sort.Ints(bansMagicNumber)
	for _, banNumber := range bansMagicNumber {
		if n >= int64(banNumber) {
			n++
		} else {
			break
		}
	}

	answer := ""
	for n > 0 {
		reminder := (n - 1) % 26
		answer = string(rune(reminder+97)) + answer
		n = (n - 1) / 26
	}
	return answer
}

func calculateMagicNumber(str string) int {
	magicNumber := 0
	for i := len(str) - 1; i >= 0; i-- {
		strRune := rune(str[i])
		alphabetNumber := int(strRune) - 96
		magicNumber += alphabetNumber * int(math.Pow(26, float64(len(str)-1-i)))
	}

	return magicNumber
}
