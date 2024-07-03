package main

/*
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/submissions/1136235176/\
처음에는 map을 이용해서 풀었다. 문제는 문제없이 풀렸지만 성능이 너무 안나와 다른 방법을 찾아봤다.
두번째 방법은 two pointer방식으로 풀었다. 첫번째 방식보다는 조금 더 빠른 속도를 보여줬다.
*/

func twoSum(numbers []int, target int) []int {

	numIndex := make(map[int]int, len(numbers))
	answer := make([]int, 2)
	for i, v := range numbers {
		if val, exist := numIndex[target-v]; exist {
			answer[0] = val + 1
			answer[1] = i + 1
		}
		numIndex[v] = i
	}

	return answer
}

func optimizedTwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	var sum int
	for left < right {
		sum = numbers[left] + numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			break
		}
	}

	return []int{left + 1, right + 1}
}
