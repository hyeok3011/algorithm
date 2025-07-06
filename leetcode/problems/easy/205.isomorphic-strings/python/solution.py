# https://leetcode.com/problems/isomorphic-strings/description/

class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        mapping = {}
        t_set = set()
        for s_char, t_char in zip(s, t):
            if s_char not in mapping:
                if t_char in t_set:
                    return False
                mapping[s_char] = t_char
                t_set.add(t_char)
            if t_char != mapping[s_char]:
                return False
        return True

# 똑똑쓰...
class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        zipped_set = set(zip(s, t))
        return len(zipped_set) == len(set(s)) == len(set(t))