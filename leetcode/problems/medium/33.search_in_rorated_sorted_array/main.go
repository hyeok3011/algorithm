package main

/*
https://leetcode.com/problems/search-in-rotated-sorted-array/
정렬된 슬라이스에서 특정 인덱스 기준으로 서로 로테이트된 상태이다.
O(n)으로 풀다면 매우매우 쉬운 문제이지만 O(log n)으로 풀려니 어떻게 풀어야할지 감이 안온다.
정렬된 상태라면 바이너리 서치를 사용하여  O(logn)으로 쉽게 풀리겠지만 위 문제는 한번 로테이트 되어있다.
내가 아는 방법중 O(logn) 방법은 이것뿐이니 먼저 시도해보자
왼쪽이든 오른쪽이든 오름차순으로 정렬되어있을것이다.
먼저 mid와 target과 비교하고 왼쪽에 있는지 오른쪽에 있는지 확인하는 식으로 풀었다.
틀린케이스보며 끼워맞춰가면서 풀어서 뭔가 찜찜하다
풀고 다른사람들의 풀이를 보니 크게 다르지 않다. 더 신선한 방법이 있을법 한데..
*/

func search(numbers []int, target int) int {
	var left, mid, right int
	left, right = 0, len(numbers)-1

	for left <= right {
		mid = left + (right-left)/2
		if numbers[mid] == target {
			return mid
		} else if numbers[left] <= numbers[mid] {
			if target < numbers[mid] && target >= numbers[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > numbers[mid] && target <= numbers[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
