#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'arrayManipulation' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts following parameters:
#  1. INTEGER n
#  2. 2D_INTEGER_ARRAY queries
#
# https://www.hackerrank.com/challenges/crush/problem?isFullScreen=true
# 효율적인 정답을 답을 찾지 못해 검색결과 아래와같은 방법이 있음. 정리가 필요함
# @@@@@@@@@@@@@@@@@@@@@@@@
# FailCode Brute Force
def arrayManipulationFail(n, queries):
    # Write your code here
    score = [0] * n
    max_val = 0
    for query in queries:
        start, end, val = query[0], query[1], query[2]
        for index in range(start-1, end):
            score[index] += val
            if max_val < score[index]:
                max_val = score[index]

    return max_val

# Prefix Sum / Difference Array
def arrayManipulation(n, queries):
    # Write your code here
    score = [0] * (n + 1)

    for query in queries:
        start, end, val = query[0], query[1], query[2]
        score[start - 1] += val
        score[end] += (val * -1)
    
    sum = 0
    max_val = 0
    for index in range(0, n):
        sum += score[index]
        if max_val < sum:
            max_val = sum

    return max_val

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    m = int(first_multiple_input[1])

    queries = []

    for _ in range(m):
        queries.append(list(map(int, input().rstrip().split())))

    result = arrayManipulation(n, queries)

    fptr.write(str(result) + '\n')

    fptr.close()
