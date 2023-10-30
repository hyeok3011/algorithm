package main

// https://school.programmers.co.kr/learn/courses/30/lessons/87389
// 먼저 n이 홀수라면 무조건 무조건 정답은 2가 될것이라 먼저 제외 시키고
// n이 짝수라면 3 5 7 9 처럼 홀수의 값으로만 확인해보면 된다.

func solution(n int) int {
	if n%2 == 1 {
		return 2
	}

	answer := 3
	for n%answer != 1 {
		answer += 2
	}
	return answer
}
