package main

// https://leetcode.com/problems/unique-number-of-occurrences/description/?envType=study-plan-v2&envId=leetcode-75
func uniqueOccurrences(arr []int) bool {
	numOccurrenceCount := make(map[int]int)
	for _, v := range arr {
		if _, exist := numOccurrenceCount[v]; !exist {
			numOccurrenceCount[v] = 0
		}
		numOccurrenceCount[v]++
	}

	occurrenceSet := make(map[int]struct{})
	for _, count := range numOccurrenceCount {
		if _, exist := occurrenceSet[count]; exist {
			return false
		}
		occurrenceSet[count] = struct{}{}
	}
	return true
}
