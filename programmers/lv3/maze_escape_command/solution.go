package main

// 기본 d, l, r, u순서대로 접근하며 제외사항을 제외하고나니 풀렸다... 사실 풀릴지 몰랐다.
// 제외사항 2가지
// 이동해야할 거리보다 이동할수 있는 거리가 짧은경우.
// 이동할수있는 거리가 남기는하지만 우회해서 가거나 탈출지점에서 움직이기 위해서는 2번은 움직여야 하기때문에 이 경우도 제외
func solution(n int, m int, x int, y int, r int, c int, k int) string {
	cmd := findPathDFS(n, m, x, y, r, c, k, "")
	if cmd == "" {
		cmd = "impossible"
	}
	return cmd
}

func findPathDFS(n int, m int, x int, y int, r int, c int, depth int, command string) string {
	distance := getManhattanDisance(x, y, r, c)
	if distance > depth || (distance-depth)%2 != 0 {
		return ""
	}
	if depth == 0 {
		return command
	}
	if x < n {
		if cmd := findPathDFS(n, m, x+1, y, r, c, depth-1, command+"d"); cmd != "" {
			return cmd
		}
	}
	if y > 1 {
		if cmd := findPathDFS(n, m, x, y-1, r, c, depth-1, command+"l"); cmd != "" {
			return cmd
		}
	}

	if y < m {
		if cmd := findPathDFS(n, m, x, y+1, r, c, depth-1, command+"r"); cmd != "" {
			return cmd
		}
	}

	if x > 1 {
		if cmd := findPathDFS(n, m, x-1, y, r, c, depth-1, command+"u"); cmd != "" {
			return cmd
		}
	}
	return ""
}

func getManhattanDisance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
