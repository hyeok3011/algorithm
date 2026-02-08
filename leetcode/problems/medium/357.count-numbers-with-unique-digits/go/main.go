package main

// n의 길이만큼 숫자 순열을 만든다고 접근한 방법
// https://leetcode.com/problems/count-numbers-with-unique-digits/solutions/7405240/beats-100-go-simple-dp-solution-on-time-evoy5/
// @@@@@@@@ 위 풀이는 생각하지 못했던방식....
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	digits := [10]bool{}

	var permutations func(current int) int
	permutations = func(current int) int {
		if current == 0 {
			return 1
		}

		count := 0
		for i := 0; i < 10; i++ {
			if digits[i] {
				continue
			}
			digits[i] = true
			count += permutations(current - 1)
			digits[i] = false
		}
		return count
	}

	result := 10
	for i := 2; i <= n; i++ {
		for d := 1; d <= 9; d++ {
			digits[d] = true
			result += permutations(i - 1)
			digits[d] = false
		}
	}

	return result
}
