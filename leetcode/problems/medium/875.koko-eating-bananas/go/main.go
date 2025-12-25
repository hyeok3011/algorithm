package main

// https://leetcode.com/problems/koko-eating-bananas/description/?envType=study-plan-v2&envId=leetcode-75
// 문제에 나와있듯이 minimum integer k  를 구하는거를 생각해야함.
func minEatingSpeed(piles []int, h int) int {
	maxBanana := 0
	for i := range piles {
		maxBanana = max(maxBanana, piles[i])
	}

	left, right := 1, maxBanana

	for left <= right {
		mid := (left + right) / 2

		if h-calculateHours(piles, mid) >= 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func calculateHours(piles []int, k int) int {
	times := 0
	for i := range piles {
		times += (piles[i] + k - 1) / k
	}

	return times
}

// 다른사람으 ㅣ코드를 보고 early return한 사용한 버전
// 다 먹을수 있느냐 없느냐에 집중
func minEatingSpeed2(piles []int, h int) int {
	maxBanana := 0
	for i := range piles {
		maxBanana = max(maxBanana, piles[i])
	}

	left, right := 1, maxBanana

	for left <= right {
		mid := (left + right) / 2

		if canEatAll(piles, h, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func canEatAll(piles []int, h, k int) bool {
	times := 0
	for i := range piles {
		times += (piles[i] + k - 1) / k
		if times > h {
			return false
		}
	}

	return true
}
