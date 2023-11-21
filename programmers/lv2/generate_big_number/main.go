package main

// https://school.programmers.co.kr/learn/courses/30/lessons/42883
// 스택을 사용하여 각 수자를 확인하고 현재 숫자가 스택의 마지막 숫자보다 큰 경우 스택에서
// 제거하는 방식으로 풀었다.
// 그리고 같은 숫자만 있는 경우를 대비해서 마지막에 스택에서 -k만큼 제거하였다.
func solution(number string, k int) string {
	stack := make([]byte, 0, len(number))

	for i := 0; i < len(number); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] < number[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, number[i])
	}

	return string(stack[:len(stack)-k])
}
