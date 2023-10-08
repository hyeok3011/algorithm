package main

// https://school.programmers.co.kr/learn/courses/30/lessons/118667
// 두 큐의 합을 동일하게 만드는 문제다.
// 각각 큐의 합을 구한뒤 greedy방식으로 문제를 풀었다.
func solution(queue1 []int, queue2 []int) int {
	firstQueueSum := sumOfQueue(queue1)
	secondQueueSum := sumOfQueue(queue2)
	step := 0
	maxStepCount := (len(queue1) + len(queue2)) * 2
	for firstQueueSum != secondQueueSum && step < maxStepCount {
		if firstQueueSum > secondQueueSum {
			firstQueueSum -= queue1[0]
			secondQueueSum += queue1[0]
			queue2 = append(queue2, queue1[0])
			queue1 = queue1[1:]
		} else {
			secondQueueSum -= queue2[0]
			firstQueueSum += queue2[0]
			queue1 = append(queue1, queue2[0])
			queue2 = queue2[1:]
		}
		step++
	}

	if firstQueueSum != secondQueueSum {
		return -1
	}

	return step
}

func sumOfQueue(queue []int) int {
	sum := 0
	for _, v := range queue {
		sum += v
	}
	return sum
}
