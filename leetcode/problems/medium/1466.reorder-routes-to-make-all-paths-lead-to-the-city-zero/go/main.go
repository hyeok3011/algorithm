package main

// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/?envType=study-plan-v2&envId=leetcode-75
func minReorder(n int, connections [][]int) int {
	inBound := make([][]int, n)
	outBound := make([][]int, n)
	for i := range connections {
		start, end := connections[i][0], connections[i][1]
		outBound[start] = append(outBound[start], end)
		inBound[end] = append(inBound[end], start)
	}

	answer := 0
	var recursion func(city, patents int)
	recursion = func(city, parents int) {
		for _, nextCity := range inBound[city] {
			if nextCity != parents {
				recursion(nextCity, city)
			}
		}

		for _, nextCity := range outBound[city] {
			if nextCity != parents {
				answer++
				recursion(nextCity, city)
			}
		}
	}

	recursion(0, 0)
	return answer
}
