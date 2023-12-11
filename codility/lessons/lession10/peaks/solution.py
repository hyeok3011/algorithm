# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

"""
https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/peak_indices/start/
배열을 동일한크기의 블록으로 나누는데 그 블록 안에는 반드시 peak가 포함되어 있어야한다.
문제 자체는 크기 어렵지 않다.
1. peak를 구한다.
2. 최대 블록의 개수는 Peak개수만큼이기 때문에 peak개수부터 차례대로 내려온다.
3. 블록안에 peak가 들어가있는지 확인한다.

"""
def solution(A):
    N = len(A)
    peak_indices = []

    for i in range(1, N - 1):
        if A[i - 1] < A[i] > A[i + 1]:
            peak_indices.append(i)

    if not peak_indices:
        return 0

    peak_blocks = set()
    for block_count in range(len(peak_indices), 0, -1):
        if N % block_count != 0:
            continue
        
        block_size = N // block_count

        for peak in peak_indices:
            peak_block_index = peak // block_size
            if peak_block_index not in peak_blocks:
                peak_blocks.add(peak_block_index)

        if len(peak_blocks) == block_count:
            break
        peak_blocks.clear()

    return len(peak_blocks)
