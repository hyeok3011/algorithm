# https://leetcode.com/problems/roman-to-integer/description/
# Runtime 54 ms Beats 83.83% Memory 16.3 MB Beats 36.27%
ROMAN = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000             
        }
# MCMXCIV 1994
class Solution:
    def romanToInt(self, s: str) -> int:
        i = len(s) - 1
        integer = 0
        last_num = 0
        while i >= 0:
            current = ROMAN[s[i]]
            if last_num <= current:
                integer += current
            else:
                integer -= current
            last_num = current
            i -= 1
        return integer
    