package main

// https://school.programmers.co.kr/learn/courses/30/lessons/161988

func solution(sequence []int) int64 {
	negativeSum := int64(0)
	positiveSum := int64(0)
	negativeMax := int64(-100000)
	positiveMax := int64(-100000)

	for i := range sequence {
		val := int64(sequence[i])
		if i%2 == 0 {
			negativeSum = max(negativeSum-val, -val)
			positiveSum = max(positiveSum+val, val)
		} else {
			negativeSum = max(negativeSum+val, val)
			positiveSum = max(positiveSum-val, -val)
		}
		negativeMax = max(negativeMax, negativeSum)
		positiveMax = max(positiveMax, positiveSum)
	}

	return max(negativeMax, positiveMax)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
