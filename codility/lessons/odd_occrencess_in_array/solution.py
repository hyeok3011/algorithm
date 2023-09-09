# https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/

def solution(A):
    # Implement your solution here
    odd_number_map = {}
    for odd_number in A:
        if odd_number in odd_number_map:
            odd_number_map.pop(odd_number)
        else:
            odd_number_map[odd_number] = True

    return list(odd_number_map.keys())[0]

# 예전에 봤던 XOR연산방법도 있다.
#  9 ^= 3의 경우
#   1001  (9 in binary)
#   0011  (3 in binary)
#   1010
# 알고는 있지만 뭔가 알고리즘 문제 풀때 딱하고 생각이 나지 않는다...
def solution2(A):
    result = 0
    for number in A:
        result ^= number
    return result
