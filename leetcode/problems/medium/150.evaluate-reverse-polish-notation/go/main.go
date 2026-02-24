package main

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for i := range tokens {
		if isOperator(tokens[i]) {
			first, second := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, calculate(first, second, tokens[i]))
		} else {
			number, _ := strconv.Atoi(tokens[i])
			stack = append(stack, number)
		}
	}

	return stack[0]
}

func isOperator(char string) bool {
	return char == "*" || char == "-" || char == "+" || char == "/"
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}
