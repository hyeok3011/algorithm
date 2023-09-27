package main

// https://school.programmers.co.kr/learn/courses/30/lessons/159994
func solution(firstDeck []string, secondDeck []string, goal []string) string {
	answer := "Yes"

	firstDeckIndex := 0
	secondDeckIndex := 0

	for _, v := range goal {
		if firstDeckIndex < len(firstDeck) && firstDeck[firstDeckIndex] == v {
			firstDeckIndex += 1
		} else if secondDeckIndex < len(secondDeck) && secondDeck[secondDeckIndex] == v {
			secondDeckIndex += 1
		} else {
			answer = "No"
			break
		}
	}

	return answer
}
