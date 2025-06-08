package main

// https://leetcode.com/problems/dota2-senate/description/?envType=study-plan-v2&envId=leetcode-75
// 문제 지문으로는 이해할수가 없음... discussion을 봐야지만 이해 가능...
func predictPartyVictory(senate string) string {
	radiantQueue := make([]int, 0, len(senate))
	direQueue := make([]int, 0, len(senate))

	for i, v := range senate {
		if v == 'R' {
			radiantQueue = append(radiantQueue, i)
		} else {
			direQueue = append(direQueue, i)
		}
	}
	lastIndex := len(senate)
	for len(radiantQueue) > 0 && len(direQueue) > 0 {
		if radiantQueue[0] < direQueue[0] {
			direQueue = direQueue[1:]
			radiantQueue = radiantQueue[1:]
			radiantQueue = append(radiantQueue, lastIndex)
		} else {
			radiantQueue = radiantQueue[1:]
			direQueue = direQueue[1:]
			direQueue = append(direQueue, lastIndex)
		}
		lastIndex++
	}

	if len(direQueue) > 0 {
		return "Dire"
	}
	return "Radiant"
}
