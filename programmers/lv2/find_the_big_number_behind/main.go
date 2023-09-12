package main

func solution(numbers []int) []int {
	numIndexStack := []int{0}

	answer := make([]int, len(numbers))

	for i := 1; i < len(numbers); i++ {
		currentVal := numbers[i]
		for len(numIndexStack) > 0 && numbers[numIndexStack[len(numIndexStack)-1]] < currentVal {
			answer[numIndexStack[len(numIndexStack)-1]] = currentVal
			numIndexStack = numIndexStack[:len(numIndexStack)-1]
		}
		numIndexStack = append(numIndexStack, i)
	}

	for i := range answer {
		if answer[i] == 0 {
			answer[i] = -1
		}
	}

	return answer
}
