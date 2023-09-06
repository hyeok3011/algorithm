# https://app.codility.com/programmers/lessons/5-prefix_sums/genomic_range_query/start/
IMPACT_FACTOR = {
    "A": 1,
    "C": 2,
    "G": 3,
    "T": 4
}
# @@@@@@@@@@@@@@@@@@@@@
# 위 링크에서 prefix sum힌트를 주지 않았다면 timeout해결방법을 찾지못해 쉽게 못풀었을거같음...
def build_frequency(S, nucleotide):
    count = 0
    frequency = [0]
    for char in S:
        if char == nucleotide:
            count += 1
        frequency.append(count)
    return frequency

def solution(S, P, Q):
    frequencies = {}
    for char in ["A", "C", "G", "T"]:
        frequencies[char] = build_frequency(S, char)
    
    minimal_impact_factors = []
    
    for index in range(0, len(Q)):
        q_val = Q[index]
        p_val = P[index]
        for nucleotide in ["A", "C", "G", "T"]:
            if frequencies[nucleotide][q_val + 1] - frequencies[nucleotide][p_val] > 0:
                minimal_impact_factors.append(IMPACT_FACTOR[nucleotide])
                break
                
    return minimal_impact_factors

solution("CAGCCTA", [2,5,0], [4,5,6])