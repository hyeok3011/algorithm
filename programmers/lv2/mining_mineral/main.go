package main

import "sort"

func solution(picks []int, minerals []string) int {
	pickFatigueValues := map[int]map[string]int{
		0: {"diamond": 1, "iron": 1, "stone": 1},
		1: {"diamond": 5, "iron": 1, "stone": 1},
		2: {"diamond": 25, "iron": 5, "stone": 1},
	}

	totalPick := 0
	for _, v := range picks {
		totalPick += v
	}

	mineralGroup := [][]int{}
	index := -1

	for i, v := range minerals {
		if i%5 == 0 {
			index += 1
			if index >= totalPick {
				break
			}
			mineralGroup = append(mineralGroup, []int{0, i})
		}
		mineralGroup[index][0] += pickFatigueValues[2][v]
	}

	sort.Slice(mineralGroup, func(i, j int) bool {
		return mineralGroup[i][0] > mineralGroup[j][0]
	})

	totalFatigue := 0

	for _, group := range mineralGroup {
		currentPick := getPickAndReduce(picks)

		startIndex := group[1]
		endIndex := startIndex + 5
		if endIndex >= len(minerals) {
			endIndex = len(minerals)
		}

		for i := startIndex; i < endIndex; i++ {
			mineral := minerals[i]
			totalFatigue += pickFatigueValues[currentPick][mineral]
		}
	}
	return totalFatigue
}

func getPickAndReduce(picks []int) int {
	pick := -1
	for i, v := range picks {
		if v != 0 {
			pick = i
			picks[i]--
			break
		}
	}
	return pick
}
