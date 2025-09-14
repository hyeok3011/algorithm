package main

const (
	MIN_INT = -2147483648
	MAX_INT = 2147483647
)

// https://leetcode.com/problems/string-to-integer-atoi/
// int64 오버플로우 케이스는 처음봄..
func myAtoi(s string) int {
	i := 0
	n := len(s)
	if n == 0 {
		return 0
	}

	for i < n && s[i] == ' ' {
		i++
	}

	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	answer := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		if answer > MAX_INT/10 || (answer == MAX_INT/10 && digit > 7) {
			if sign == 1 {
				return MAX_INT
			}
			return MIN_INT
		}

		answer = answer*10 + digit
		i++
	}

	return answer * sign
}
