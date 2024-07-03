package main

// https://leetcode.com/problems/excel-sheet-column-number/description/
// Runtime 0 ms Beats 100% Memory 2.1 MB Beats 98.69%
func titleToNumber(columnTitle string) int {
	len := len(columnTitle) - 1
	answer := 0
	for _, v := range columnTitle {
		sum := v - 64
		for i := 0; i < len; i++ {
			sum *= 26
		}

		answer += int(sum)
		len -= 1
	}
	return answer
}

func titleToNumber2(columnTitle string) int {
	answer := 0

	for _, char := range columnTitle {
		digit := int(char - 'A' + 1)
		answer = (answer * 26) + digit
	}

	return answer
}
