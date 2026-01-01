package main

import "container/heap"

// https://leetcode.com/problems/smallest-number-in-infinite-set/description/
// 사용할때마다 heap에 숫자를 하나씩 넣어주는 방식에서 최적화 진행
type SmallestInfiniteSet struct {
	set     map[int]struct{}
	minHeap IntHeap
	head    int
}

func Constructor() SmallestInfiniteSet {
	h := IntHeap{}
	heap.Init(&h)
	intSet := make(map[int]struct{})
	return SmallestInfiniteSet{
		set:     intSet,
		minHeap: h,
		head:    1,
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if s.minHeap.Len() > 0 {
		top := heap.Pop(&s.minHeap).(int)
		delete(s.set, top)
		return top
	}
	v := s.head
	s.head++
	return v
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := s.set[num]; ok {
		return
	}

	if num >= s.head {
		return
	}

	heap.Push(&s.minHeap, num)
	s.set[num] = struct{}{}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
