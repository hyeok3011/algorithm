package main

// https://leetcode.com/problems/number-of-provinces/description/
func findCircleNum(isConnected [][]int) int {
	provinces := make([]int, len(isConnected))

	var recursion func(startCircle, targetCircle int)

	recursion = func(startCircle, targetCircle int) {
		if provinces[targetCircle] != 0 {
			return
		}
		provinces[targetCircle] = startCircle
		for i := range isConnected[targetCircle] {
			if isConnected[targetCircle][i] == 0 {
				continue
			}
			recursion(startCircle, i)
		}
	}

	for i := range isConnected {
		recursion(i+1, i)
	}

	set := make(map[int]struct{})

	for i := 0; i < len(provinces); i++ {
		set[provinces[i]] = struct{}{}
	}

	return len(set)
}
