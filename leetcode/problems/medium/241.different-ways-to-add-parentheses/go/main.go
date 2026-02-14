package main

// https://leetcode.com/problems/different-ways-to-add-parentheses/
// @@@@ 생각보다 구현이 어려웠음. 다시 풀어봐야함...
func diffWaysToCompute(expression string) []int {
	numbers := []int{}
	operations := []byte{}
	number := 0
	for i := range expression {
		char := expression[i]
		if '0' <= char && char <= '9' {
			number *= 10
			number += int(char - '0')
		} else {
			numbers = append(numbers, number)
			number = 0
			operations = append(operations, char)
		}
	}
	numbers = append(numbers, number)

	dp := make([][]int, len(numbers))
	for i := range dp {
		dp[i] = make([]int, len(numbers))
	}

	memo := make([][][]int, len(numbers))
	for i := range memo {
		memo[i] = make([][]int, len(numbers))
	}

	var recursion func(i, j int) []int
	recursion = func(i, j int) []int {
		if memo[i][j] != nil {
			return memo[i][j]
		}
		if i == j {
			memo[i][j] = []int{numbers[i]}
			return memo[i][j]
		}

		result := []int{}
		for k := i; k < j; k++ {
			left := recursion(i, k)
			right := recursion(k+1, j)
			for l := range left {
				for r := range right {
					result = append(result, calculator(left[l], right[r], operations[k]))
				}
			}
		}
		memo[i][j] = result
		return result
	}

	return recursion(0, len(numbers)-1)
}

func calculator(a, b int, operation byte) int {
	if operation == '-' {
		return a - b
	} else if operation == '+' {
		return a + b
	} else {
		return a * b
	}
}
