package main

// https://leetcode.com/problems/fibonacci-number/description/
// Runtime 0 ms Beats 100% Memory 1.9 MB Beats 99.40%
func fib(n int) int {
	visited := make(map[int]int)
	return fibonacciNumber(n, visited)
}

func fibonacciNumber(n int, visited map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	if v, ok := visited[n]; ok {
		return v
	}

	visited[n] = fibonacciNumber(n-1, visited) + fibonacciNumber(n-2, visited)

	return visited[n]
}
