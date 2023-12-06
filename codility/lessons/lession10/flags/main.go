package main

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

/*
https://app.codility.com/c/run/trainingDMTYZ6-QTF/
먼저 peak부터 모두 확인 후 최대 꼽을 수 있는 깃발의 개수를 확인해야한다.
먼저 확인해야할 범위는 총 길이의 제곱근까지만 가능하다 이유는 거리 + 깃발 개수라는 조건이 있기 때문이다.
어... 틀린 케이스에서 sqrt + 1이 있는것을 보아 다른 예외 케이스가 있는것 같은데 어떤 상황인지 모르겠다...
*/
func Solution(A []int) int {
	// Implement your solution here
	peaks := []int{}
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	if len(peaks) == 0 {
		return 0
	}

	maxFlagCount := 0
	sqrt := sqrt(len(A)) + 1
	for i := 1; i <= sqrt; i++ {
		if validPlaceFlag(peaks, i) {
			maxFlagCount = max(maxFlagCount, i)
		}
	}
	return maxFlagCount
}
func sqrt(x int) int {
	left, right := 0, x

	for left <= right {
		mid := left + (right-left)/2
		midSquared := mid * mid
		if midSquared > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func validPlaceFlag(peaks []int, flag int) bool {
	prePeak := peaks[0]
	count := 1
	for i := 1; i < len(peaks); i++ {
		if peaks[i]-prePeak >= flag {
			count += 1
			prePeak = peaks[i]

			if count == flag {
				return true
			}
		}
	}

	return count == flag
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
