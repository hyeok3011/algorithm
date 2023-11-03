import math
import os
import random
import re
import sys
from collections import Counter

#
# Complete the 'matchingStrings' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts following parameters:
#  1. STRING_ARRAY stringList
#  2. STRING_ARRAY queries
#

def matchingStrings(stringList, queries):
    # Write your code here
    recordMap = {}
    for data in stringList:
        if data in recordMap:
            recordMap[data] += 1
        else:
            recordMap[data] = 1
            
    matched_count = []
    for i in range(0, len(queries)):
        if queries[i] in recordMap:
            matched_count.append(recordMap[queries[i]])
        else:
            matched_count.append(0)
    
    return matched_count

# god python
def matchingStrings(stringList, queries):
    # Write your code here
    recordMap = Counter(stringList)
            
    return [recordMap.get(query, 0) for query in queries]

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    stringList_count = int(input().strip())

    stringList = []

    for _ in range(stringList_count):
        stringList_item = input()
        stringList.append(stringList_item)

    queries_count = int(input().strip())

    queries = []

    for _ in range(queries_count):
        queries_item = input()
        queries.append(queries_item)

    res = matchingStrings(stringList, queries)

    fptr.write('\n'.join(map(str, res)))
    fptr.write('\n')

    fptr.close()
