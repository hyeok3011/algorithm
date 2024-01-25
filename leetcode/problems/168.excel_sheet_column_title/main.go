package main

/*
https://leetcode.com/problems/excel-sheet-column-title/description/
1 = A 26 = Z 28 = AB
숫자가 주어지면 해당 숫자를 문자로 바꾸는 문제이다.
*/

func convertToTitle(columnNumber int) string {
	answer := ""
	remainder := 0
	for columnNumber > 0 {
		remainder = columnNumber % 26
		columnNumber = columnNumber / 26
		if remainder == 0 {
			remainder = 26
			columnNumber--
		}
		answer = string(64+rune(remainder)) + answer
	}

	return answer
}
