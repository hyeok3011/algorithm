package main

import "container/heap"

type item struct {
	value int
	count int
}

// 정렬 - (n + m log m)
// Heap - (n + m log k) k개 유지
func topKFrequent(nums []int, k int) []int {
	numCount := make(map[int]int)
	for _, v := range nums {
		numCount[v]++
	}
	intHeap := &IntHeap{}
	heap.Init(intHeap)
	for num, count := range numCount {
		if intHeap.Len() < k {
			heap.Push(intHeap, item{
				value: num,
				count: count,
			})
		} else if (*intHeap)[0].count < count {
			heap.Pop(intHeap)
			heap.Push(intHeap, item{
				value: num,
				count: count,
			})
		}
	}

	result := make([]int, k)
	for i, item := range *intHeap {
		result[i] = item.value
	}
	return result
}

type IntHeap []item

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(item))
}

func (h *IntHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

// 다른 사람의 풀이.... Bucket Sort 진짜 생각도 못했음.
// @@@@
// topKFrequent finds the k most frequent elements in an array.
// Example: nums = [1,1,1,2,2,3], k = 2 returns [1,2] as 1 occurs 3 times and 2 occurs 2 times
// Time Complexity: O(n) where n is the length of nums
// Space Complexity: O(n) for storing counts and frequency array
func topKFrequentBucketSort(nums []int, k int) []int {
	// Step 1: Create a map to count frequency of each number
	// Example: [1,1,1,2,2,3] -> map[1:3, 2:2, 3:1]
	count := make(map[int]int)

	// Step 2: Create an array where index represents frequency-1
	// frequency[i] will store all numbers that appear i+1 times
	// Example: frequency[2] will store numbers that appear 3 times
	frequency := make([][]int, len(nums))

	// Step 3: Count occurrences of each number
	// Iterate through input array and increment count for each number
	for _, num := range nums {
		count[num] += 1
	}

	// Step 4: Build frequency array
	// For each number and its count in our map
	// Add the number to the array at index (count-1)
	// Example: if 1 appears 3 times, add 1 to frequency[2]
	for num, freq := range count {
		frequency[freq-1] = append(frequency[freq-1], num)
	}

	// Step 5: Build result array by taking most frequent elements
	res := []int{}
	// Start from the end of frequency array (most frequent elements)
	// and work backwards until we have k elements
	for i := len(frequency) - 1; i >= 0; i-- {
		// For each number that appears i+1 times
		for _, num := range frequency[i] {
			res = append(res, num)
			// If we have found k elements, return result
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
