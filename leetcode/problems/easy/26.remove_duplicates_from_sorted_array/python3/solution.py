# https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
# Runtime 94 ms Beats 86.59% Memory 18 MB Beats 18.50%
class Solution:
    def removeDuplicates(self, nums) -> int:
        pre = 0
        for index in range (1, len(nums)):
            if nums[pre] == nums[index]:
                continue
            pre += 1
            nums[pre] = nums[index]
        return pre + 1