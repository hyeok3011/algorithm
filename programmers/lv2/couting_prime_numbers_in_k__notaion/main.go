package main

// https://school.programmers.co.kr/learn/courses/30/lessons/92335
// k진수를 만드는 방법은 number / k의 나머지를 구하고 reverse시키면 된다. (주입된 지식)
// 그 다음부터는 간단하다
// 소수판별은 i * i이상까지는 할 필요가 없다.
// 그 수의 제곱근보다 큰 두 숫자의 곱으로 나누어 떨어지지 않기 때문에 제곱근까지만 검사해도 충분하다
// 그리고 그 수가 2로 나누어진다면 일단 소수가 아닐것이고 3부터 홀수만 검사하면 된다. (주입된 지식)
func solution(n int, k int) int {
	number := 0
	numberPlace := 1
	primeCount := 0
	for n > 0 {
		remainder := n % k
		n = n / k

		if remainder == 0 {
			if isPrime(number) {
				primeCount++
			}
			number = 0
			numberPlace = 1
		} else {
			number += remainder * numberPlace
			numberPlace *= 10
		}
	}

	if isPrime(number) {
		primeCount++
	}

	return primeCount
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	if number%2 == 0 {
		return false
	}

	for i := 3; i*i <= number; i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}
