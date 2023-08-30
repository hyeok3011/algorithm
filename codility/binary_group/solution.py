# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(N):
    # Implement your solution here
    binary = bin(N)[3:]
    
    print(binary)
    max_count = 0
    count = 0
    for v in binary:
        if v == "0":
            count += 1
        else:
            if max_count < count:
                max_count = count
            count = 0
    return max_count

def solution2(N):
    binary = bin(N)[2:]
    zero_sequences = binary.split("1")

    max_count = max([len(sequence) for sequence in zero_sequences])
    return max_count
