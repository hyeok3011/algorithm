package main

// https://leetcode.com/problems/sqrtx/description/
// Runtime 0 ms Beats 100% Memory 2.1 MB Beats 99.64%
// midSqured와 x를 비교 안하고 right를 반환해도 문제는 없는 이유는
// right가 최종적으로 x의 제곱근이 되고 left가 더 커지기 때문
// 기존 무식한 방법으로 for i := 0; i < x/2+1 방식으로 하면 너무 오래걸린다.
func mySqrt(x int) int {
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
