"""
https://leetcode.com/problems/restore-ip-addresses/description/
쉬운문제였으나 깔끔하게 풀기 어려웠음...@@@@@@@@@@@@
"""

from typing import List

class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        def is_valid(segment: str) -> bool:
            return 0 <= int(segment) <= 255 and (segment == "0" or segment[0] != "0")
        
        def backtrack(start: int, path: List[str]):
            if len(path) == 4:
                if start == len(s):
                    result.append(".".join(path))
                return
            
            for length in range(1, 4):
                if start + length <= len(s):
                    segment = s[start:start + length]
                    if is_valid(segment):
                        backtrack(start + length, path + [segment])
        
        result = []
        backtrack(0, [])
        return result