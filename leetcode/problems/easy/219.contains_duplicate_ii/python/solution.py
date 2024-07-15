# https://leetcode.com/problems/contains-duplicate-ii/description/

# 단순한 hash table문제이다. 코드가 420ms정도 나오는데 70ms까지 나오는 코드가 있는듯 한데 어떻게 가능한건지 참...
from typing import List

class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        num_index_map = {}
        for i, num in enumerate(nums):
            if num in num_index_map and (i - num_index_map[num]) <= k:
                return True
            num_index_map[num] = i
        return False
