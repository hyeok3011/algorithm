package main

// https://leetcode.com/problems/gas-station/description/
// @@@@@@ 처음에 O(n^2)으로 풀었으나 시간초과 발생.
// i ~ j 사이의 skip을 통해 O(n)으로 풀었으나 생각하는 사고 과정이 매끄럽지 못함.
func canCompleteCircuit(gas []int, cost []int) int {
	totalSum := 0
	currentSum := 0
	startIndex := 0
	for i := range gas {
		totalSum += (gas[i] - cost[i])
		currentSum += (gas[i] - cost[i])
		if currentSum < 0 {
			currentSum = 0
			startIndex = i + 1
		}
	}

	if totalSum < 0 {
		return -1
	}

	return startIndex
}
