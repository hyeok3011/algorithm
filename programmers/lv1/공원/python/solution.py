# https://school.programmers.co.kr/learn/courses/30/lessons/340198
# 개선 버전. all()좋은것 같음, 전체 for을 사용하지 않고 matsize만큼 빼서 loop
def is_placable(park, r_start, c_start, size):
    return all(
        park[r][c] == "-1"
        for r in range(r_start, r_start + size)
        for c in range(c_start, c_start + size)
    )

def solution(mat_sizes, park):
    mat_sizes.sort(reverse=True)

    for size in mat_sizes:
        for r in range(len(park) - size + 1):
            for c in range(len(park[0]) - size + 1):
                
                if is_placable(park, r, c, size):
                    return size 

    return -1