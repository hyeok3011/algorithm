package main

// type Stack struct {
// 	elements []rune
// }

// func (s *Stack) Push(v rune) {
// 	s.elements = append(s.elements, v)
// }

// func (s *Stack) Top() rune {
// 	return s.elements[len(s.elements)-1]
// }

// func (s *Stack) isEmpty() bool {
// 	if len(s.elements) == 0 {
// 		return true
// 	}
// 	return false
// }

// func (s *Stack) Pop() rune {
// 	v := s.Top()
// 	s.elements = s.elements[:len(s.elements)-1]
// 	return v
// }

// func (s *Stack) Clear() {
// 	s.elements = []rune{}
// }

// func (s *Stack) Len() int {
// 	return len(s.elements)
// }

// func fail(s string) int {
// 	var validParentheses, maxValidParentheses int
// 	stack := Stack{}
// 	for _, char := range s {
// 		if char == '(' {
// 			stack.Push(char)
// 		} else {
// 			if !stack.isEmpty() {
// 				stack.Pop()
// 				validParentheses += 2
// 			} else {
// 				maxValidParentheses = max(maxValidParentheses, validParentheses)
// 				validParentheses = 0
// 				stack.Clear()
// 			}
// 		}
// 	}
// 	maxValidParentheses = max(maxValidParentheses, validParentheses)

// 	return maxValidParentheses
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
