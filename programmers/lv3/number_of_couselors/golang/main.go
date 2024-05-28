package main

import (
	"container/heap"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/214288
// 기존에 slcie를 이용하여 풀었는데 정답오류가 계속 나서 우선순위큐로 바꾸었더니 갑자기 성공함... 뭐지?
func solution(k int, n int, reqs [][]int) int {

	mentorCountByType := make(map[int]int)

	requestsByType := make(map[int][][]int)
	for i := 0; i < k; i++ {
		requestsByType[i] = [][]int{}
		mentorCountByType[i] = 1
	}

	for _, v := range reqs {
		startTime, duration, reqType := v[0], v[1], v[2]-1
		requestsByType[reqType] = append(requestsByType[reqType], []int{startTime, duration})
	}

	waitTimes := make([]int, len(mentorCountByType))
	for i := 0; i < k; i++ {
		waitTimes[i] = computeWaitTime(requestsByType[i], mentorCountByType[i])
	}

	for n-k > 0 {
		maxEfficenceIndex := 0
		maxEfficence := 0
		for i := 0; i < len(mentorCountByType); i++ {
			waitTime := computeWaitTime(requestsByType[i], mentorCountByType[i]+1)
			if waitTimes[i]-waitTime > maxEfficence {
				maxEfficenceIndex = i
				maxEfficence = waitTimes[i] - waitTime
			}
		}

		mentorCountByType[maxEfficenceIndex] += 1
		waitTimes[maxEfficenceIndex] = waitTimes[maxEfficenceIndex] - maxEfficence
		n -= 1
	}

	sum := 0
	for _, v := range waitTimes {
		sum += v
	}

	return sum
}

func computeWaitTime(requests [][]int, mentorsCount int) int {
	waitTime := 0

	minHeap := &IntHeap{}
	heap.Init(minHeap)
	for i := 0; i < mentorsCount; i++ {
		heap.Push(minHeap, 0)
	}

	for _, req := range requests {
		startTime, duration := req[0], req[1]
		earliestMentor := heap.Pop(minHeap).(int)

		if earliestMentor > startTime {
			waitTime += earliestMentor - startTime
			heap.Push(minHeap, earliestMentor+duration)
		} else {
			heap.Push(minHeap, startTime+duration)
		}
	}

	return waitTime
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

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
