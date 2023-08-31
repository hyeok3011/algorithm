# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")
def solution(S: str):
    # Implement your solution here
    splited = S.split(" ")
    max_length = 0
    is_one_word = False
    if len(splited) == 1:
        splited = split_by_non_alphanumeric(splited[0])
        is_one_word = True


    for word in splited:
        if is_valid_password(word):
            max_length = max(max_length, len(word))

    if is_one_word and max_length == 0 and not is_alphanumeric(S):
        max_length = -1

    return max_length

def split_by_non_alphanumeric(text: str) -> list:
    word = ""
    alphanumerices = []
    for char in text:
        if is_alphanumeric(char):
            word += char
        else:
            if not word.isspace():
                alphanumerices.append(word)
    return alphanumerices
    
def is_alphanumeric(text: str) -> bool:
    for char in text:
        if not char.isdigit() and not char.isalpha():
            return False
    return True

def is_valid_password(text: str) -> bool:
    letter_count = 0
    digit_count = 0
    
    for v in text:
        if v.isdigit():
            digit_count += 1
        elif v.isalpha():
            letter_count += 1
        else:
            return False

    if letter_count % 2 == 1 or letter_count == 0:
        return False
    
    if digit_count % 2 == 0 or digit_count == 0:
        return False

    return True

