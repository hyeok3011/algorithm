package main

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
		if apeachScoreCopy < lionScore && diff >= maxDiff {
			if diff == maxDiff && isGreater(lionInfo, maxDiffInfo) {
				maxDiff = diff
				maxDiffInfo = lionInfo
			} else {
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

func isGreater(a, b []int) bool {
	for i := 10; i >= 0; i-- {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	return false
}
