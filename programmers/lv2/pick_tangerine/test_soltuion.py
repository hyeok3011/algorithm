from solution import solution

def test_solution():
    assert solution(6, [1, 3, 2, 5, 4, 5, 2, 3]) == 3
    assert solution(4, [1, 3, 2, 5, 4, 5, 2, 3]) == 2
    assert solution(2, [1, 1, 1, 1, 2, 2, 2, 3]) == 1