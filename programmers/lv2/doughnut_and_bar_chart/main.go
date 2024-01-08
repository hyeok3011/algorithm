package main

/*
https://school.programmers.co.kr/learn/courses/30/lessons/258711
처음에는 그래프들을 모두 순회하며 문제를 해결하려고 했다. 하지만 위 방법으로 풀다가 ㅈㅈ첬다.
결국 질문하기에서 힌트를 조금 얻었다.
결국은 패턴을 활용하여 푸는것이 정답이라는것
결국 막대 경우에는 보내는 엣지가 없어야 한다.
그리고 8모양의 경우에는 중간 정점기준 2개를 보내고 2개 이상을 받을 수 있다.
도넛 모양은 사실 어떤 패턴이 있나 엄청 고민해봤지마ㅣㄴ 발견하지 못하고 total - bar - 8개수를 빼니까 정답이 나왔다.
root에서 나간 정점중에 새로운 차트가 생기는 상화잉였다면 풀리지 않았을텐데 다행이다.

root정점은 받는것은 없고 보내는것만 있으며 아래의 문제 예시를 따라 확인하였다.
도넛 모양 그래프, 막대 모양 그래프, 8자 모양 그래프의 수의 합은 2이상입니다.
*/

func solution(edges [][]int) []int {

	sendEdgeCount := make(map[int]int)
	receiveEdgeCount := make(map[int]int)
	lastVertexNumber := 0
	for _, edge := range edges {
		sendEdgeCount[edge[0]]++
		receiveEdgeCount[edge[1]]++
		lastVertexNumber = max(edge[0], edge[1], lastVertexNumber)
	}

	var rootVertex, doughnutCount, barCount, eightCount int
	for i := 1; i <= lastVertexNumber; i++ {
		if receiveEdgeCount[i] == 0 && sendEdgeCount[i] >= 2 {
			rootVertex = i
		}

		if sendEdgeCount[i] == 0 {
			barCount++
		} else if sendEdgeCount[i] == 2 && receiveEdgeCount[i] >= 2 {
			eightCount++
		}
	}

	doughnutCount = sendEdgeCount[rootVertex] - barCount - eightCount
	return []int{rootVertex, doughnutCount, barCount, eightCount}
}

func max(numbers ...int) int {
	maxVal := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if maxVal < numbers[i] {
			maxVal = numbers[i]
		}
	}
	return maxVal
}
