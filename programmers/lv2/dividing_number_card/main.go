package main

// https://school.programmers.co.kr/learn/courses/30/lessons/135807
// 영희와 철수 카드중 본인 카드는 모두 나눠지고 상대방 카드는 모두 나눠지지 않는 숫자중 가장 큰 숫자를 찾는 문제이다.
// 먼저 본인 카드의 최대 공약수를 찾아야한다는 것을 알 수 있다.
// 그리고 상대방 카드에서 본인의 최대 공약수가 약수가 되는지만 확인해봤더니 모두 풀렸다.
// 문제를 조금더 꼬았다면 복잡해젔을거같다.
func solution(arrayA []int, arrayB []int) int {
	var yh, cs int
	if gcd := getGCD(arrayA); gcd != 1 {
		if isNotCD(arrayB, gcd) {
			yh = gcd
		}
	}

	if gcd := getGCD(arrayB); gcd != 1 {
		if isNotCD(arrayA, gcd) {
			cs = gcd
		}
	}

	if yh > cs {
		return yh
	}

	return cs
}

func isNotCD(elements []int, gcd int) bool {
	for _, v := range elements {
		if v%gcd == 0 {
			return false
		}
	}

	return true
}

func getGCD(elements []int) int {
	gcd := elements[0]

	for _, elem := range elements[1:] {
		gcd = euClidGCD(gcd, elem)
	}

	return gcd
}

func euClidGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
