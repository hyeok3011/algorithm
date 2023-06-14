# Runtime 26 ms Beats 98.38% Memory 16.2 MB Beats 77.96%
# https://leetcode.com/problems/length-of-last-word/description/
class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        return len(s.split()[-1])
