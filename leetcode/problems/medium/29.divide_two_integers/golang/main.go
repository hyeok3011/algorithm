package main

// https: //leetcode.com/problems/divide-two-integers/description/
// optimized the code with the help of chatgpt. The optimization will be comment on.
// Runtime 0 ms Beats 100% Memory 2.3 MB Beats 98.98%
const (
	MAXIMUM = 1<<31 - 1
	MINIMUM = -1 << 31
)

func divide(dividend int, divisor int) int {
	signal := getSignal(dividend) * getSignal(divisor)
	dividend = abs(dividend)
	divisor = abs(divisor)
	result := 0

	for dividend >= divisor {
		shifts := 0
		for dividend >= (divisor << shifts) {
			shifts++
		}
		shifts--
		dividend -= divisor << shifts
		result += 1 << shifts
	}

	result *= signal

	if result < MINIMUM {
		return MINIMUM
	}

	if result > MAXIMUM {
		return MAXIMUM
	}

	return result
}

func getSignal(a int) int {
	if a >= 0 {
		return 1
	}
	return -1
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

// https://leetcode.com/problems/divide-two-integers/description/
// Runtime 0 ms Beats 100% Memory 2.3 MB Beats 39.34%
func divide2(dividend int, divisor int) int {
	result := int(dividend / divisor)
	if MINIMUM > result {
		return MINIMUM
	}

	if MAXIMUM < result {
		return MAXIMUM
	}
	return result
}
