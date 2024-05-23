# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")
def solution(A, B, C):
    # 먼저 2진수 문자열로 변환
    result_set = set()
    for i in range(0, 30):
        val = 1 << i
        a_bit = A & val
        b_bit = B & val
        c_bit = C & val

        if (a_bit and b_bit and c_bit) or (not a_bit and not b_bit and not c_bit):
            continue
        

        result_set.add(A if a_bit else A + val)
        result_set.add(B if b_bit else B + val)
        result_set.add(C if c_bit else C + val)

    return len(result_set)

def solution(A, B, C):
    a_bit_slice =int_to_bit_slice()

def to_padded_binary_str(num, length=30):
    bin_str = bin(num)[2:] 
    bin_slice = bin_str.zfill(length)
    return bin.sl

if __name__ == "__main__":
    print(solution(1073741727, 1073741631, 1073741679))