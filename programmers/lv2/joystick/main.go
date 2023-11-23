package main

// https://school.programmers.co.kr/learn/courses/30/lessons/42860?language=
// 문제 카테고리가 탐욕방법인데 도저히 안풀려서 사람들이 어떻게 풀었나 보니 많은사람들이  완전탐색으로
// 푸는것 보고 포기할뻔했으나 간신히 풀었다.
// 알파벳을 맞추기 위해 움직이는 조이스틱의 횟수는 정해저있다.
// 그래서 죄우로 어떻게 움직일지에 대해서 고민했다.
// 좌우로 움직이는 경우는
// 1. 우측으로 쭉 움직이는 경우
// 2. 우측으로 가다가 좌측으로 다시 움직이는 경우
// 3. 좌측으로 가다가 다시 돌아와서 우측으로 가는경우

// 1번 케이스는 우측으로 쭉 움직이는 상황이니 n - 1이다.
// 2,3번은 현재 위치에서 연속되는 A가 끝나는 부분을 확인후 왼오, 오왼을 계산하며 최솟값을 찾는다.

func solution(name string) int {
	n := len(name)
	minMove := n - 1

	total := 0
	for i := 0; i < n; i++ {
		total += calculateDistanceFromA(name[i])

		nextNonA := i + 1
		for nextNonA < n && name[nextNonA] == 'A' {
			nextNonA++
		}

		// 오른쪼이동 후 왼쪽
		minMove = min(minMove, i*2+n-nextNonA)
		// 왼쪽 이동 후 오른쪽
		minMove = min(minMove, (n-nextNonA)*2+i)
	}

	return total + minMove
}

func calculateDistanceFromA(char byte) int {
	return min(int(char-'A'), int('Z'-char+1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
