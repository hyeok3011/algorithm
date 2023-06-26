
# https://leetcode.com/problems/valid-perfect-square/description/
# Runtime 44 ms Beats 71.80% Memory 16.2 MB Beats 62.81%
# binary search algorithm
class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        if num == 1:
            return True

        left = 2
        right = num
        
        while left <= right:
            mid = left + (right - left) // 2
            mid_squared = mid * mid

            if mid_squared == num:
                return True
            elif mid_squared < num:
                left = mid + 1
            else:
                right = mid - 1
        
        return False
