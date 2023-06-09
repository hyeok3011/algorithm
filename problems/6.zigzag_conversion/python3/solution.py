# https://leetcode.com/problems/zigzag-conversion/description/
# Runtime 60 ms Beats 85.74% Memory 16.4 MB Beats 41.11%
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1 or numRows >= len(s):
            return s
        
        result = [''] * numRows
        index, step = 0, 1
        
        for char in s:
            result[index] += char
            
            if index == 0:
                step = 1
            elif index == numRows - 1:
                step = -1
            
            index += step
        
        return ''.join(result)
    
# Runtime 68 ms Beats 69.3% Memory 16.5 MB Beats 41.11%
class Solution2:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        answer = ""
        numRows -= 1 
        for num_row in range(0, numRows+1):
            index = num_row
            up = True
            while index < len(s):
                answer += s[index]
                if index % numRows == 0:
                    index += numRows + numRows
                elif up:
                    index += (numRows - num_row) * 2 
                elif not up:
                    index +=  num_row * 2
                up = not up
                    
        return answer