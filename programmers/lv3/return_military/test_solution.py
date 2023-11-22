from solution import solution

def test_solution():
    assert solution(3, [[1, 2], [2, 3]], [2,3], 1) == [1, 2]
    assert solution(5, [[1, 2], [1, 4], [2, 4], [2, 5], [4, 5]], [1, 3, 5], 5) == [2, -1, 0]