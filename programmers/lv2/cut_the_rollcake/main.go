package main

// https://school.programmers.co.kr/learn/courses/30/lessons/132265
// 어떻게 풀지 처음에 고민을 했다.
// 2중 for문을 돌리면 분명 시간초과가 뜰거기 때문이다.
// 그래서 생각한것이 처음에 영희에게 모든 토핑을 몰아주고 다시 순회하며 영희꺼를 하나씩 뺏어보는
// 방식으로 풀었더니 문제없이 풀렸다.
func solution(topping []int) int {
	csToppingSet := make(map[int]int)
	yhToppingSet := make(map[int]int)
	for i := range topping {
		yhToppingSet[topping[i]] += 1
	}

	answer := 0
	for i := range topping {
		csToppingSet[topping[i]] += 1
		yhToppingSet[topping[i]] -= 1

		if yhToppingSet[topping[i]] == 0 {
			delete(yhToppingSet, topping[i])
		}

		if len(csToppingSet) == len(yhToppingSet) {
			answer += 1
		}
	}

	return answer
}

// 다른 사람의 풀이법이다.
// 토핑의 개수 맥시멈만큼 미리 토핑의 종류 개수 배열을 철수 영희꺼를 모두 만들어두고
// 포팅을 순회하며 철수는 왼쪽부터, 영희는 오른쪽부터 토핑의 종류 개수 배열을 채워나간다.
// 두번째 순회에서는 두 배열의 해당 인덱스에서 같은 토핑 개수가 나오는 부분을 확인하는 방법이다.
// 생각하지 못한 방법이다. 하지만 토핑개수 맥시멈만큼 배열을 만드는게 조금 아쉽다.
func solution2(topping []int) int {
	answer := 0
	dp := [2][1000001]int{}
	mL := map[int]bool{}
	mR := map[int]bool{}

	for i := 0; i < len(topping); i++ {
		// L
		mL[topping[i]] = true
		dp[0][i] = len(mL)

		// R
		mR[topping[len(topping)-i-1]] = true
		dp[1][len(topping)-i-1] = len(mR)
	}

	for i := 0; i < len(topping)-1; i++ {
		if dp[0][i] == dp[1][i+1] {
			answer++
		}
	}

	return answer
}
