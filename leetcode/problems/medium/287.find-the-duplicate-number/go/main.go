package main

// https://leetcode.com/problems/find-the-duplicate-number/description/
// @@@@@@@@@@ 
func findDuplicateSet(nums []int) int {
    set := make(map[int]bool, len(nums))
    for _, v := range nums {
        if set[v] {
            return v
        }
        set[v] = true
    }

    return -1
}

// Floyd으로 푸는사람 발견... 중복이 존재하고 싸이클이 반드시 생기기때문에 가능한것같음.
func findDuplicate(nums []int) int {
    slow := nums[0]
    fast := nums[nums[0]]
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }

    slow = 0
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }

    return slow
}
