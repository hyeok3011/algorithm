package main

import (
	"sort"
)

/*
https://school.programmers.co.kr/learn/courses/30/lessons/258709
보자마자 쉽게 풀리지 않을 문제라는 느낌이 왔다.
1. 주사위 선택의 조합 경우의수
2. 조합볍 6 ** n만큼의 조합 합 경우의수
주사위가 10까지 나온경우 답도 없겠다라는 생각이 들었다.

여기서 든 생각은
1. 캐시를 활용하여 비교 연산부분을 최대한 줄여보자
2. 플레이어 들끼리의 주사위 비교 연산 자체를 효율적으로 해보자.
먼저 가장 간단한 1번부터 시도해본다.
음.... 시간초과로 조금 실패한다. 캐시로 활용가능한 부분이 더 많을거 같지만 찾이 못했다.
그러면 비교연산을 조금 더효율적인 방법을 생각해보자, 현재는 O(n^2)
단순 순회가 아닌 정렬후 투 포인터 방법을 사용하는 방법??? 을 새악ㄱ해보자
만약 {1, 1, 80, 80, 80, 80},{1, 1, 70, 70, 70, 70} 을 비교했을때
위 1 1 은 뒤로 갈 필요도 없다.
근데 중요한것은 정렬또한 시간을 잡아 먹는다는것 golang의 정렬은 O(n*log(n)
그러면 (2 * nLogn) + (nlogn)이 된다...
충분히 더 빠를수있겠다라는 생각이 든다.
시도해보자

*/

func solution(dice [][]int) []int {
	combinations := [][]int{}
	generateDiceCombinations(0, 0, len(dice), make([]int, len(dice)/2), &combinations)
	maxWinCount := 0
	bestDiceCombination := []int{}

	for _, playerADiceIndexes := range combinations {
		var playerAWinCount int

		playerBDiceIndexes := findOtherIndexes(playerADiceIndexes, len(dice))
		playerADiceSumCombinations := []int{}
		playerBDiceSumCombinations := []int{}
		diceSumCombination(0, playerADiceIndexes, dice, 0, &playerADiceSumCombinations)
		diceSumCombination(0, playerBDiceIndexes, dice, 0, &playerBDiceSumCombinations)

		sort.Ints(playerADiceSumCombinations)
		sort.Ints(playerBDiceSumCombinations)

		for _, aScore := range playerADiceSumCombinations {
			index := sort.Search(len(playerBDiceSumCombinations), func(i int) bool {
				return playerBDiceSumCombinations[i] >= aScore
			})
			playerAWinCount += index
		}

		if maxWinCount < playerAWinCount {
			maxWinCount = playerAWinCount
			bestDiceCombination = playerADiceIndexes
		}
	}

	answer := make([]int, len(bestDiceCombination))
	for i := range bestDiceCombination {
		answer[i] = bestDiceCombination[i] + 1
	}

	return answer
}

func diceSumCombination(index int, indexes []int, dice [][]int, currentSum int, combinations *[]int) {
	if index == len(indexes) {
		*combinations = append(*combinations, currentSum)
		return
	}

	for i := 0; i < 6; i++ {
		currentSum += dice[indexes[index]][i]
		diceSumCombination(index+1, indexes, dice, currentSum, combinations)
		currentSum -= dice[indexes[index]][i]
	}
}

func findOtherIndexes(indexes []int, total int) []int {
	indexMap := make(map[int]bool, len(indexes))
	for _, v := range indexes {
		indexMap[v] = true
	}
	otherIndexes := []int{}
	for i := 0; i < total; i++ {
		if !indexMap[i] {
			otherIndexes = append(otherIndexes, i)
		}
	}

	return otherIndexes
}

func generateDiceCombinations(depth, startIndex, n int, current []int, combinations *[][]int) {
	if depth == n/2 {
		temp := make([]int, len(current))
		copy(temp, current)
		*combinations = append(*combinations, temp)
		return
	}

	for i := startIndex; i < n; i++ {
		current[depth] = i
		generateDiceCombinations(depth+1, i+1, n, current, combinations)
	}
}
