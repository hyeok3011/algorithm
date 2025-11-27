// https://leetcode.com/problems/task-scheduler/description/
func leastInterval(tasks []byte, n int) int {
	counts := make(map[byte]int)
	for _, t := range tasks {
		counts[t]++
	}

	pq := &IntHeap{}
	heap.Init(pq)
	for _, c := range counts {
		heap.Push(pq, c)
	}

	answer := 0

	for pq.Len() > 0 {
		var temp []int
		window := n + 1

		for pq.Len() > 0 && window > 0{
			count := heap.Pop(pq).(int)
			
			count--
			
			if count > 0 {
				temp = append(temp, count)
			}
			
			answer++
			window--
		}

		for _, remainingCount := range temp {
			heap.Push(pq, remainingCount)
		}

		if pq.Len() > 0 {
			answer += window
		}
	}

	return answer
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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


// 다른사람의 풀이......
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	cnt := make([]int, 26)
	for _, task := range tasks {
		cnt[task - 'A']++
	}

	var maxCount, sameMaxCount int
	for _, c := range cnt {
		if c > maxCount {
			maxCount = c
			sameMaxCount = 1
		} else if c == maxCount {
			sameMaxCount++
		}
	}

	res := (n + 1) * (maxCount - 1) + sameMaxCount
    if (res > len(tasks)) {
        return res
    } else {
        return len(tasks)
    }
}
