from solution import Solution

s = Solution()

def test_maxProfit():
    assert s.maxProfit([7,1,5,3,6,4]) == 7
    assert s.maxProfit([1,2,3,4,5]) == 4
    assert s.maxProfit([7,6,4,3,1]) == 0