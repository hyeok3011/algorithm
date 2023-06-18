package main

const (
	MAXIMUM = 2<<31 - 1
	MINIMUM = -2 << 31
)

// https://leetcode.com/problems/integer-to-roman/
// Runtime 7 ms Beats 81.66% Memory 3.3 MB Beats 38.75%
func intToRoman(num int) string {
	integers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	answer := ""
	index := len(roman) - 1
	for index >= 0 {
		if num >= integers[index] {
			answer += roman[index]
			num -= integers[index]
		} else {
			index -= 1
		}
	}
	return answer
}
