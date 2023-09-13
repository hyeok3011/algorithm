package main

// https://school.programmers.co.kr/learn/courses/30/lessons/178870
func solution(sequence []int, k int) []int {
	result := []int{0, len(sequence) - 1}
	start := 0
	end := 0
	sum := sequence[0]
	for start <= end {
		if sum < k && end < len(sequence)-1 {
			end += 1
			sum += sequence[end]
		} else {
			if sum == k && (result[1]-result[0]) > end-start {
				result[0] = start
				result[1] = end
			}
			sum -= sequence[start]
			start += 1
		}
	}

	return result
}
