package main

// https://school.programmers.co.kr/learn/courses/30/lessons/92342
// 지금까지 문제중 최악의 문제였다. 조건이 너무나도 까다로웠다... 이런 문제는 너무 싫다....
// 라이언이 어피치와 가장 큰 점수차이를 내며 이기는 케이스를 확인하는 문제이다.
// 조건은
// 같은 점수의 과녁을 맞추면 어피치가 점수를 얻게되고 라이언은 반드시 높아야한다.
// 가장 큰 점수차이를 만들더라도 더 낮은 점수의 비중이 높아야한다.

func solution(n int, info []int) []int {
	apeachScore := 0
	for i := range info {
		if info[i] > 0 {
			apeachScore += 10 - i
		}
	}

	maxDiff := 0
	maxDiffInfo := make([]int, len(info))
	scores := [][]int{}
	generateCombination(n, 0, info, make([]int, len(info)), &scores)

	for _, lionInfo := range scores {
		lionScore := 0
		apeachScoreCopy := apeachScore
		for j := range lionInfo {
			if info[j] < lionInfo[j] {
				lionScore += 10 - j
				if info[j] != 0 {
					apeachScoreCopy -= 10 - j
				}
			}
		}

		diff := lionScore - apeachScoreCopy
		if diff > maxDiff {
			maxDiff = diff
			maxDiffInfo = lionInfo
		} else if diff == maxDiff {
			if hasHigherLowScoreWeight(lionInfo, maxDiffInfo) {
				maxDiff = diff
				maxDiffInfo = lionInfo
			}
		}
	}
	if maxDiff == 0 {
		return []int{-1}
	}
	return maxDiffInfo
}

func generateCombination(n, start int, info []int, current []int, combinations *[][]int) {
	if start == 11 {
		combination := make([]int, len(info))
		if n == 0 {
			copy(combination, current)
		} else {
			current[start-1] = n
			copy(combination, current)
		}
		*combinations = append(*combinations, combination)
		return
	}

	if n >= info[start]+1 {
		current[start] = info[start] + 1
		generateCombination(n-current[start], start+1, info, current, combinations)
		current[start] = 0
	}
	generateCombination(n, start+1, info, current, combinations)
	current[start] = 0
	return
}

func hasHigherLowScoreWeight(a, b []int) bool {
	for i := 10; i >= 0; i-- {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	return false
}
