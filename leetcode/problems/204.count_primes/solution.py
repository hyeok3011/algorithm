"""
https://leetcode.com/problems/count-primes/
숫자 n이 들어오면 n까지 몇개의 소수가 들어있는지 확인하는 문제이다.
다양한 방법을 시도해봤지만 모두 타임아웃으로 실패한다... 이런문제 너무 싫다...
역시 방법은 에라토스테네스의 체 뿐이 없는듯 싶다....
https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
"""

class Solution:
    def countPrimes(self, n: int) -> int:
        if n < 3:
            return 0

        is_prime = [True] * n
        is_prime[0],is_prime[1] = False, False

        for i in range(2, int(n**0.5) + 1):
            if is_prime[i]:
                for j in range(i*i, n, i):
                    is_prime[j] = False

        return sum(is_prime)
