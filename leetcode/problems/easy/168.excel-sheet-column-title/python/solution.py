# https://leetcode.com/problems/excel-sheet-column-title/description/
class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        chars = []
        while columnNumber > 0:
            columnNumber -= 1
            chars.append(chr(ord('A') + columnNumber % 26))
            columnNumber = columnNumber // 26
        return ''.join(reversed(chars))
