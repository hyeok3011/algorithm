def solution(rows, columns, queries):
    matrix = []
    count = 1
    for i in range(rows):
        row = []
        for j in range(columns):
            row.append(count)
            count+=1
        matrix.append(row)

    
    answer = []
    for x1, y1, x2, y2 in queries:
        min_val = rotation_and_find_min(matrix, x1-1, y1-1, x2-1, y2-1) 
        answer.append(min_val)
    return answer

def rotation_and_find_min(matrix, x1, y1, x2, y2) -> int:
    temp = matrix[x1][y1]
    min_val = temp
    
    for i in range(x1, x2):
        matrix[i][y1] = matrix[i+1][y1]
        min_val = min(min_val, matrix[i][y1])

    for i in range(y1, y2):
        matrix[x2][i] = matrix[x2][i+1]
        min_val = min(min_val, matrix[x2][i])
    
    for i in range(x2, x1, -1):
        matrix[i][y2] = matrix[i-1][y2]
        min_val = min(min_val, matrix[i][y2])
    
    for i in range(y2, y1+1, -1):
        matrix[x1][i] = matrix[x1][i-1]
        min_val = min(min_val, matrix[x1][i])
    
    matrix[x1][y1+1] = temp
    return min_val