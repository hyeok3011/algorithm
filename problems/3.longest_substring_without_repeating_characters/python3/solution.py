# https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
# Runtime 89ms Beats 31.78% Memory 16.3MB Beats 61.75%
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) < 2:
            return len(s)
        start, end = 0, 0
        substring = set()
        max_count = 0
        while end < len(s):
            if s[end] not in substring:
                substring.add(s[end])
                end += 1
                max_count = max(max_count, len(substring))
            else:
                substring.remove(s[start])
                start += 1
            
        return max_count

