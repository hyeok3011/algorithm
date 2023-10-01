package main

// https://school.programmers.co.kr/learn/courses/30/lessons/147355
func solution(t string, p string) int {
	smallerCount := 0
	pLen := len(p)
	index := 0
	for index+pLen <= len(t) {
		substring := t[index : index+pLen]
		if substring <= p {
			smallerCount++
		}
		index++
	}

	return smallerCount
}
