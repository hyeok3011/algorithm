from solution import solution, solution2

def test_solution_per():
    assert solution([4,1,3,2]) == 1
    
def test_solution_not_per():
    assert solution([4,1,3]) == 0
    
def test_solution2_per():
    assert solution2([4,1,3,2]) == 1
    
def test_solution2_not_per():
    assert solution2([4,1,3]) == 0