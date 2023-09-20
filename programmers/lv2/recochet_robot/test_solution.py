from solution import solution

def test_solution():
    assert solution(["...D..R", ".D.G...", "....D.D", "D....D.", "..D...."]) == 7

def test_solution_fail():
    assert solution([".D.R", "....", ".G..", "...D"]) == -1
