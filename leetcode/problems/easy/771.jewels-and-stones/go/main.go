package main

// https://leetcode.com/problems/jewels-and-stones/description/
// Array를 사용하면 Hashig비용을 아끼기때문에 더 효율적으로 풀수는 있을거같음.
func numJewelsInStones(jewels string, stones string) int {
	jewleSet := make(map[byte]bool, len(jewels))
	for i := range jewels {
		jewleSet[jewels[i]] = true
	}

	count := 0
	for i := range stones {
		if jewleSet[stones[i]] {
			count++
		}
	}

	return count
}
