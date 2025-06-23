# 
def solution(n, w, num):
    box_id = num - 1

    target_row = box_id // w
    pos_in_row = box_id % w

    if target_row % 2 == 0:
        target_col = pos_in_row
    else:
        target_col = w - 1 - pos_in_row

    max_row_idx = (n - 1) // w
    
    top_box_in_col_num = 0
    if max_row_idx % 2 == 0:
        top_box_in_col_num = (max_row_idx * w) + target_col + 1
    else:
        pos_from_right = w - 1 - target_col
        top_box_in_col_num = (max_row_idx * w) + pos_from_right + 1
        
    stack_height = 0
    if top_box_in_col_num > n:
        stack_height = max_row_idx
    else:
        stack_height = max_row_idx + 1

    answer = stack_height - target_row
    
    return answer
