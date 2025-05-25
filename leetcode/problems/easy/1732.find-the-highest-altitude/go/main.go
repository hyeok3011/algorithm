package main

// https://leetcode.com/problems/find-the-highest-altitude/?envType=study-plan-v2&envId=leetcode-75
func largestAltitude(gain []int) int {
	currentAltitude := 0
	higherAltitude := 0
	for _, v := range gain {
		currentAltitude += v
		higherAltitude = max(higherAltitude, currentAltitude)
	}

	return higherAltitude
}
