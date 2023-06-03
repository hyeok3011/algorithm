
# Runtime 47ms Beats 43.17% Memory 16.2MB Beats 60.1%
class Solution:
    def reverse(self, x: int) -> int:
        sign = 1
        if x < 0:
            sign = -1
            x *= sign

        result = 0
        while x != 0:
            result = result*10 + x%10
            x = int(x/10)

        if result > 1<<31:
            return 0
        
        return result * sign
    
    def reverse2(self, x: int) -> int:
        sign = 1
        result = 0
        if x < 0:
            sign = -1
        result = int(str(x*sign)[::-1])
        if 1<<31 < result:
            return 0
        return result *sign

if __name__ == "__main__":
    sol = Solution()
    print(sol.reverse(-123))