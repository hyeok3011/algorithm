"""
https://leetcode.com/problems/power-of-two/description/
수ㅅ자 n이 주어지면 n이 2의 거듭제곱인지 확인하는 문제이다.
간단하게 recursion이나 loop로 풀면 되겠다 싶었지만 두개 사용 금지라고 한다...
먼저 n % 2 == 1이 되는 경우는 제외이다.
그래서 생각한것이 2의 거듭제급의 이진수는 항상 1이 한개만 존재한다.
그래서 이진수로 변환 수 1의 개수를 확인하면 쉽게 풀수 있을듯 하다.
"""


class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        return n > 0 and bin(n).count("1") == 1