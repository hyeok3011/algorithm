# Runtime 46 ms Beats 61.69% Memory 16.2 MB Beats 70.49%
# https://leetcode.com/problems/string-to-integer-atoi/
MAX = 2 ** 31 - 1
MIN = -2 ** 31

class Solution:
    def myAtoi(self, s: str) -> int:
        integer = 0
        sign = 1
        s = s.lstrip()
        if len(s) == 0:
            return 0

        if self.is_str_signal(s[0]):
            sign = self.string_signal_to_int(s[0])
            s = s[1:]

        for char in s:
            if char.isdigit():
                integer = integer * 10 + int(char)
            else:
                break
        integer = sign * integer
        if MAX <= integer:
            return MAX
        if MIN >= integer:
            return MIN

        return integer
    
    def is_str_signal(self, char: str) -> bool:
        return "+" == char or "-" == char
    
    def string_signal_to_int(self, char: str) -> int:
        if char == "-":
            return -1
        return 1