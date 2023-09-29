package main

// https://school.programmers.co.kr/learn/courses/30/lessons/150369
// 집  #1  #2 #3  #4 #5
// 배달	1개	0개	3개	1개	2개
// 수거	0개	3개	0개	4개	0개
// 생각보다 매우 어려웠던 문제 처음에는 문제의 예시처럼 뒤에서부터 탐색하며 배달개수 및 픽업개수를 확인 후
// 이전에 있는것들까지 모두 탐색하는 식으로 문제를 풀었으나 매우 복잡한 방식이였음.
// 이후 다른사람들의 풀이를 보니 그리디 방식으로 매우 쉽게 풀리는것을 보고 충격먹음...
// @@@@@@@@@@@@@@@@@@@@ 다시볼것
func solution(cap int, n int, deliveries []int, pickups []int) int64 {
	var moveDistance int64
	var deliveryCount, pickupCount int

	for i := len(deliveries) - 1; i >= 0; i-- {
		if deliveries[i] == 0 && pickups[i] == 0 {
			continue
		}

		deliveryCount += deliveries[i]
		pickupCount += pickups[i]

		for deliveryCount > 0 || pickupCount > 0 {
			deliveryCount -= cap
			pickupCount -= cap
			moveDistance += int64((i + 1) * 2)
		}

	}
	return moveDistance
}
