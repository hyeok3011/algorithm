from solution import solution

def test_solution():
    assert solution([3, 4, 3, 2, 3, -1, 3, 3]) == 6
    
def test_solution_none():
    assert solution([4,  2,  -1, 3, 3]) == -1