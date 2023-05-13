package main

// https://leetcode.com/problems/two-sum/
//

func twoSum(nums []int, target int) []int {
	numberOfIndex := make(map[int]int, len(nums))
	for index, number := range nums {
		unknown := target - number
		if v, exist := numberOfIndex[unknown]; exist {
			return []int{v, index}
		}
		numberOfIndex[number] = index
	}
	return []int{}
}
