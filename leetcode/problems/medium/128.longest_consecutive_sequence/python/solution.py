"""
https://leetcode.com/problems/longest-consecutive-sequence/description/
나는 단순하게 해당 값이 시퀀스의 최대값인경우 아래값의 개수를 확인하는 방법으로 풀었으나
더 좋은 풀이방법이 있었다.
"""


from typing import List, Set

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        num_set: Set[int] = set(nums)
        longest_sequence: int = 0

        for num in nums:
            if num + 1 not in num_set:
                current_num: int = num
                current_sequence: int = 0

                while current_num in num_set:
                    current_num -= 1
                    current_sequence += 1

                longest_sequence = max(longest_sequence, current_sequence)

        return longest_sequence
    
class AnotherSolution:
    def longestConsecutive(self, nums: List[int]) -> int:
        nums = set(nums)
        table = {}
        maxval = 0
        for num in nums:
            x = table.get(num - 1, 0)
            y = table.get(num + 1, 0)
            val = x + y + 1
            table[num - x] = val
            table[num + y] = val
            maxval = max(maxval, val)
        return maxval
