# https://leetcode.com/problems/range-sum-query-immutable/
class NumArray:
    def __init__(self, nums: List[int]):
        self.nums_prefix_sum = self._calculate_pre_sum(nums)
    
    def _calculate_pre_sum(self, nums):
        pre_sum = 0
        nums_prefix_sum = [0]
        for i, v in enumerate(nums):
            pre_sum += v
            nums_prefix_sum.append(pre_sum)
        return nums_prefix_sum

    def sumRange(self, left: int, right: int) -> int:
        range_sum = self.nums_prefix_sum[right + 1]
        range_sum -= self.nums_prefix_sum[left]
        return range_sum