package main

func solution(n int, wires [][]int) int {
	minDiff := n
	parent := make([]int, n+1)
	for i := range wires {
		for j := range parent {
			parent[j] = j
		}

		for j, wire := range wires {
			if i == j {
				continue
			}
			union(parent, wire[0], wire[1])
		}

		pivotRoot := find(parent, 1)
		count := 0
		for i := 1; i < len(parent); i++ {
			if find(parent, i) == pivotRoot {
				count++
			}
		}
		minDiff = min(minDiff, abs(n-(count*2)))
	}

	return minDiff
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent []int, a, b int) {
	a = find(parent, a)
	b = find(parent, b)
	parent[a] = b
}
func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
