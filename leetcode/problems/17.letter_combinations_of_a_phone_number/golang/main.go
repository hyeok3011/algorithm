package main

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
// Runtime 0 ms Beats 100% Memory 2.1 MB Beats 56.68%

var DIGIT_LETTER = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digit := [][]string{}
	for _, v := range digits {
		digit = append(digit, DIGIT_LETTER[string(v)])
	}

	return combination(digit)
}

func combination(source [][]string) []string {
	cases := []string{}

	if len(source) == 1 {
		return source[0]
	}

	for _, val := range source[0] {
		for _, v := range combination(source[1:]) {
			cases = append(cases, val+v)
		}
	}

	return cases
}
