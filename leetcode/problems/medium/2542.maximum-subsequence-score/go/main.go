package main

// https://leetcode.com/problems/maximum-subsequence-score/?envType=study-plan-v2&envId=leetcode-75
// @@@@@ nums2는 min값을 써야함. 기준점을 정해놓고 정렬 후 처리하는 생각을 해내야함.
type Pair struct {
	n1 int
	n2 int
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
    paires := make([]Pair, len(nums1))
    for i := range paires {
        paires[i] = Pair{n1: nums1[i], n2: nums2[i]}
    }

    sort.Slice(paires, func(i, j int) bool {
		return paires[i].n2 > paires[j].n2
	})

    minHeap := IntHeap{}
    heap.Init(&minHeap)

    var maxScore int
    var sum int
    for _, pair := range paires {
        sum += pair.n1
        heap.Push(&minHeap, pair.n1)

        if minHeap.Len() > k {
            sum -= heap.Pop(&minHeap).(int)
        }

        if minHeap.Len() == k {
            maxScore = max(maxScore, sum * pair.n2)
        }
    }
    return int64(maxScore)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // MinHeap (작은게 위로)
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
