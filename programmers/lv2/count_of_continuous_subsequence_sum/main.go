package main

// https://school.programmers.co.kr/learn/courses/30/lessons/131701
// 연속 부분수열의 합의 개수를 구하는 문제이다. 여기서 수열은 원형 수열이다.
// 먼저 elements는 선형수열 이기 때문에 elements를 2배로 만들어주고
// 완전 탐색으로 풀어봤다. 분명 완전 탐색으로 풀게되면
// 첫번째 길이별 탐색 n,  수열탐색 n, 합계산 n 으로 시간복잡도는 O(n^3)이 된다.
// 분명 통과는 하는데 더 좋은 방법이 있을듯하다.
func solution(elements []int) int {
	subSequenceSet := make(map[int]bool)
	doubleElements := elements
	doubleElements = append(doubleElements, elements...)

	elemLen := len(elements)
	for i := 1; i < elemLen; i++ {
		for j := 0; j <= elemLen; j++ {
			sum := intSliceSum(doubleElements[j : j+i])
			subSequenceSet[sum] = true
		}
	}

	subSequenceSet[intSliceSum(elements)] = true
	return len(subSequenceSet)
}

func intSliceSum(intSlice []int) int {
	sum := 0
	for _, v := range intSlice {
		sum += v
	}
	return sum
}

// 아래 방법은 누군가의 도움글을 보고 풀었다.
// 방법은 이렇다.
// 2단계 부터는 1단계의 결과값과 elements에 padding을 줘 계산하는 방식이다.
// 예를들어 3단계의 경우
// 16, 10, 2, 5, 11
//
//	7,  9, 1, 1,  4, 7, 9, 1, 1, 4
//
// 에서 16 + 1, 10 + 1, 2 + 4, 5 + 7, 11 + 9
// 식으로 계산하면 문제가 풀린다.
// 이런 패턴을 어떻게 찾았는지 대단하다...
func solutionDP(elements []int) int {
	elemLen := len(elements)
	dp := make([]int, elemLen)
	copy(dp, elements)
	elements = append(elements, elements...)

	answer := make(map[int]bool)
	for _, v := range dp {
		answer[v] = true
	}

	for i := 1; i < elemLen; i++ {
		for j := 0; j < elemLen; j++ {
			dp[j] = dp[j] + elements[j+i]
			answer[dp[j]] = true
		}
	}

	return len(answer)
}
