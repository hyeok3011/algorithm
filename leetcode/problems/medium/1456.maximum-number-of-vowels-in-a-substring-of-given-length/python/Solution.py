# https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/?envType=study-plan-v2&envId=leetcode-75 
# 흠... 때려죽여도 0ms대로 안내려오는데 0ms에 았는 사람들은 어떻게 짠건지 모르겠음.
class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        vowel_letters = set(["a", "e", "i", "o", "u"])
        max_vowel_count = 0
        vowel_count = 0

        for i in range(0, k):
            if s[i] in vowel_letters:
                vowel_count += 1
        max_vowel_count = vowel_count

        for i in range(k, len(s)):
            vowel_count += (s[i] in vowel_letters) - (s[i-k] in vowel_letters)
            if max_vowel_count < vowel_count:
                max_vowel_count = vowel_count
            
        return max_vowel_count