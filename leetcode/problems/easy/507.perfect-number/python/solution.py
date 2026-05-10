# https://leetcode.com/problems/perfect-number/
class Solution:
    def checkPerfectNumber(self, num: int) -> bool:
        if num <= 1:
            return False
        
        total = 1
        for i in range(2, isqrt(num) + 1):
            if num % i == 0:
                total += i
                if i != num // i:
                    total += num // i
        return total == num

    # ??????
    def euclidEuler(self, num: int) -> bool:
        return num in {6, 28, 496, 8128, 33550336}