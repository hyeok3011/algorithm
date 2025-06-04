
# https://leetcode.com/problems/n-th-tribonacci-number/?envType=study-plan-v2&envId=leetcode-75
class Solution:
    def tribonacci(self, n: int) -> int:
        if n == 0:
            return 0
        elif n <= 2:
            return 1

        answer: Int = 0
        numbers: List[int] = [0] * (n + 1)
        numbers[0], numbers[1], numbers[2] = 0, 1, 1
        for i in range(3, n+1):
            numbers[i] = numbers[i-1] + numbers[i-2] + numbers[i-3]

        return numbers[n]

# 공간복잡도를 조금 더 줄인 코드
class Solution:
    def tribonacci(self, n: int) -> int:
        numbers: List[Int] = [0, 1, 1]
        if n < 3:
            return numbers[n]

        for i in range(3, n+1):
            numbers[0], numbers[1], numbers[2] = numbers[1], numbers[2], sum(numbers)
        return numbers[2]

# 문제에서는 분명 공간복잡도에 대한 이야기는 없었는데 다른사람들은 재귀로 많이 푼듯하다.
class Solution:
    def tribonacci(self, n: int) -> int:
        @cache
        def t(n):
            if n == 0:
                return 0
            elif n == 1 or n == 2:
                return 1
            return t(n-1) + t(n-2) + t(n-3)
        return t(n)