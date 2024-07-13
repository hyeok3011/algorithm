# https://leetcode.com/problems/contains-duplicate/description/
from typing import List

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        number_set = set()
        for num in nums:
            if num in number_set:
                return True
            number_set.add(num)
        return False
