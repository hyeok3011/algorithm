package main

// https://school.programmers.co.kr/learn/courses/30/lessons/12985
// 경기가 진행될수록 팀은 절반으로 줄어들기 때문에 a,b둘다 라운드를 진행할수록 숫자는 반으로 줄어든다.
// 1 > 1 , 2 > 1,4 = 2 홀수 같은 경우에는 강제로 짝수로 만들어줘 절반으로 줄이기 쉽게 연산한다.
// 둘이 만나기 위해서는 12, 34 이렇게 숫자가 페어가 되어야한다. 하지만 그렇게 비교하면 더욱 복잡해진다.
// 그래서 둘이 만난 뒤에는 동일한 숫자가 될것이기 때문에 만난 다음 라운드를 기준으로 정답을 확인한다.
func solution(n int, a int, b int) int {
	answer := 0

	for a != b {
		a, b = moveToNextRound(a), moveToNextRound(b)
		answer++
	}

	return answer
}

func moveToNextRound(position int) int {
	if position%2 == 0 {
		return position / 2
	}
	return (position + 1) / 2
}
