package main

// https://leetcode.com/problems/total-cost-to-hire-k-workers/description/
func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)

	leftHeap := &IntHeap{}
	rightHeap := &IntHeap{}
	heap.Init(leftHeap)
	heap.Init(rightHeap)

	leftIndex := 0
	for leftIndex < candidates && leftIndex < n {
		heap.Push(leftHeap, costs[leftIndex])
		leftIndex++
	}

	rightIndex := n - 1
	for rightIndex >= n-candidates && rightIndex >= leftIndex {
		heap.Push(rightHeap, costs[rightIndex])
		rightIndex--
	}
	var totalCost int64 = 0
	for i := 0; i < k; i++ {
		if leftHeap.Len() == 0 || (rightHeap.Len() > 0 && (leftHeap.Peek() > rightHeap.Peek())) {
			totalCost += int64(rightHeap.Peek())
			heap.Pop(rightHeap)
			if leftIndex <= rightIndex {
				heap.Push(rightHeap, costs[rightIndex])
				rightIndex--
			}
		} else {
			totalCost += int64(leftHeap.Peek())
			heap.Pop(leftHeap)
			if leftIndex <= rightIndex {
				heap.Push(leftHeap, costs[leftIndex])
				leftIndex++
			}
		}

	}
	return totalCost
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Peek() int          { return (h)[0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
