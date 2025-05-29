package main

// https://leetcode.com/problems/find-the-difference-of-two-arrays/description/?envType=study-plan-v2&envId=leetcode-75
func findDifference(numsOne []int, numsTwo []int) [][]int {
	numsOneSet := make(map[int]bool, len(numsOne))
	numsTwoSet := make(map[int]bool, len(numsTwo))
	for _, v := range numsOne {
		numsOneSet[v] = true
	}
	for _, v := range numsTwo {
		numsTwoSet[v] = true
	}

	answer := make([][]int, 2)
	answer[0] = make([]int, 0, len(numsOne))
	answer[1] = make([]int, 0, len(numsTwo))
	for v := range numsOneSet {
		if !numsTwoSet[v] {
			answer[0] = append(answer[0], v)
		}
	}
	for v := range numsTwoSet {
		if !numsOneSet[v] {
			answer[1] = append(answer[1], v)
		}
	}

	return answer
}
