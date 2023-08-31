from solution import solution

def test_solution():
    assert solution("test 5 a0A pass007 ?xy1") == 7

def test_solution_one_word():
    assert solution("aa132?asdqweqweasd1") == 5

def test_solution_one_word_and_invalid():
    assert solution("aa?asdqweqweasd1") == -1

def test_solution_one_char():
    assert solution("a") == 0


def test_solution_one_char2():
    assert solution("a1") == 0