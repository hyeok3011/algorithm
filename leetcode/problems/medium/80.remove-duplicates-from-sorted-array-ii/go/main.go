package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
func removeDuplicates(nums []int) int {
    lastSeen := 0
    count := 0
    deleted := 0

    for i, v := range nums {
        if v == lastSeen {
            count++
            if count > 2 {
                deleted++
            }
        } else {
            lastSeen = v
            count = 1
        }
        
        nums[i-deleted] = v
    }

    return len(nums) - deleted
}

// @@@@@@@ 이런 방법을.....
func removeDuplicatesOptimal(nums []int) int {
    k := 0
    for i := range nums {
        if k < 2 || nums[i] > nums[k-2] {
            nums[k] = nums[i]
            k++
        }
    }
    return k
}
