package main

/*
https://leetcode.com/problems/subsets/description/
처음에는 괄호를 스택에 쌓아놓고 다른 비슷한 문제와 같은 해결방법으로 풀어보려고 했으나
결국 실패했다.
이유는 중간에 ()(()처럼 이어지지 않거나 )(((((()())()()))()(()))(처럼 시작점을 제대로 알 수 없기 때문이다.
그래서 생각한 생각한 방법이 괄호의 인덱스 자체를 저장하여 길이를 구하는 방법을 생각했다.
인덱스로 풀며 고민됐던부분이 ()()같은 부분이다.
그래서 해결 방법으로 )나 (()에서 첫번째 (같이 의미없는 (도 구분을 위해 스택에 넣었다.
스택의 마지막이 (라면 pop하고 i - top()를 계산하면 된다.
그래도 문제가 있다.
()케이스를 틀린다. 그래서 그냥 처음에 -1을 추가해버렸다.
*/

type Stack struct {
	elements []int
}

func (s *Stack) Push(v int) {
	s.elements = append(s.elements, v)
}

func (s *Stack) Top() int {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Pop() int {
	if !s.isEmpty() {
		index := s.Top()
		s.elements = s.elements[:len(s.elements)-1]
		return index
	}
	return -1
}

func longestValidParentheses(s string) int {
	stack := Stack{}
	stack.Push(-1)
	maxLen := 0

	for i, char := range s {
		if char == '(' {
			stack.Push(i)
		} else {
			if stack.Top() != -1 && s[stack.Top()] == '(' {
				stack.Pop()
				maxLen = max(maxLen, i-stack.Top())
			} else {
				stack.Push(i)
			}
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
