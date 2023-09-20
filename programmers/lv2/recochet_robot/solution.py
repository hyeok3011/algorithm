def solution(board):
    current_position = find_start_position(board)
    queue = []
    visited = set()
    queue.append(current_position)
    
    min_depth = -1
    current_depth = 0
    while len(queue) > 0 :
        current_depth_task_len = len(queue)
        for i in range(0, current_depth_task_len):
            current_position = queue[i]
            
            if current_position in visited:
                continue
            else: 
                visited.add(current_position)
            
            x, y = current_position
            if board[x][y] == "G":
                min_depth = current_depth
                queue = queue[:current_depth_task_len]
                break
                
            queue.append(next_top_position(current_position, board))
            queue.append(next_bottom_position(current_position, board))
            queue.append(next_left_position(current_position, board))
            queue.append(next_right_position(current_position, board))
        
        queue = queue[current_depth_task_len:]
        current_depth += 1
    
    return min_depth

def find_start_position(board: list) -> tuple:
    for i in range(len(board)):
        if "R" in board[i]:
            return i, board[i].index("R")
        
def next_top_position(current_position: tuple, board: list) -> tuple:
    x, y = current_position
    while x > -1:
        if board[x][y] == "D":
            break
        x -= 1
        
    return (x + 1, y)

def next_bottom_position(current_position: tuple, board: list) -> tuple:
    x, y = current_position
    while x < len(board):
        if board[x][y] == "D":
            break
        x += 1
            
    return (x - 1, y)

def next_left_position(current_position: tuple, board: list) -> tuple:
    x, y = current_position
    while y > -1:
        if board[x][y] == "D":
            break
        y -= 1
    return (x, y + 1)

def next_right_position(current_position: tuple, board: list) -> tuple:
    x, y = current_position
    while y < len(board[0]):
        if board[x][y] == "D":
            break
        y += 1
        
    return (x, y - 1)
