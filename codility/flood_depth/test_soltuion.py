from solution import solution, solution2


def test_solution_basic():
    assert solution([1, 3, 2, 1, 2, 1, 5, 3, 3, 4, 2]) == 2

def test_solution_cannot_hold():
    assert solution([2, 5]) == 0


def test_solution2_basic():
    assert solution2([1, 3, 2, 1, 2, 1, 5, 3, 3, 4, 2]) == 2

def test_solution2_cannot_hold():
    assert solution2([2, 5]) == 0