# https://leetcode.com/problems/remove-k-digits/
class Solution:
    def removeKdigits(self, num: str, k: int) -> str:
        if len(num) == k:
            return "0"

        stack = []
        for digit in num:
            while stack and k > 0 and stack[-1] > digit:
                stack.pop()
                k -= 1
            if not stack and digit == '0':
                continue
            stack.append(digit)

        return ''.join(stack[:len(stack) - k]) or "0"