from soltuion import solution,solution2

def test_solution():
    assert solution([0, 1,0 ,1,1]) == 5
    
def test_solution_first_west():
    assert solution([1,0,0,0,0]) == 0
    
def test_solution_last_east():
    assert solution([1, 1, 1, 1 ,0]) == 0
    
def test_solution2():
    assert solution2([0, 1,0 ,1,1]) == 5
    
def test_solution2_first_west():
    assert solution2([1,0,0,0,0]) == 0
    
def test_solution2_last_east():
    assert solution2([1, 1, 1, 1 ,0]) == 0