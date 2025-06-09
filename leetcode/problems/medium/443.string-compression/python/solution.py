# https://leetcode.com/problems/string-compression/?envType=study-plan-v2&envId=leetcode-75
from typing import List


class Solution:
    def compress(self, chars: List[str]) -> int:
        write_index = 0
        current_char = chars[0]
        char_count = 1
        i = 1
        
        chars.append("#dummy#")
        
        while i < len(chars):
            if current_char == chars[i]:
                char_count += 1
            else:
                chars[write_index] = current_char
                write_index += 1
                if char_count > 1:
                    for digit in str(char_count):
                        chars[write_index] = digit
                        write_index += 1
                
                current_char = chars[i]
                char_count = 1
            i += 1
        
        return write_index


# dummy안넣고 처리
class Solution:
    def compress(self, chars: List[str]) -> int:
        write_index = 0
        current_char = chars[0]
        char_count = 1
        i = 1
        while i < len(chars):
            while i < len(chars) and chars[i] == current_char:
                char_count += 1
                i += 1
            
            chars[write_index] = current_char
            write_index += 1
            if char_count > 1:
                for digit in str(char_count):
                    chars[write_index] = digit
                    write_index += 1
            current_char = chars[i]
            char_count = 1
            i += 1
        return write_index