package main

import "sort"

// https://leetcode.com/problems/merge-sorted-array/description/
// Runtime 0 ms Beats 100% Memory 2.3 MB Beats 16.63%
// nums1의 슬라이스의 길이를 변경하지 않으면서 값을 변경해주며 0을 지워야하는데 너무 귀찮아서
// 뒤에서 값만 바꾸고 정렬을 해버렸다

// 추가로 다시 풀다보니 merge함수 내부에서는 이미 nums1의 메모리 공간을 확보해놓은 상태이기
// 때문에 nums[:m]식이로 잘라내고 append(n)을 하더라도 문제가 없는듯 하다.
func merge(nums1 []int, m int, nums2 []int, n int) {
	mergeCount := 0

	for _, v := range nums2 {
		nums1[len(nums1)-1-mergeCount] = v
		mergeCount += 1
	}

	sort.Ints(nums1)
}
