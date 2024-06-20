# 

def solution(k, dungeons):
    visited = [False for _ in dungeons]
    max_dungeon_count = 0
    def dfs(fatigue, dungeon_count):
        nonlocal max_dungeon_count
        max_dungeon_count = max(dungeon_count, max_dungeon_count)
        
        for i in range(0, len(dungeons)):
            if visited[i]:
                continue
            
            if fatigue >= dungeons[i][0] and fatigue - dungeons[i][1] >= 0:
                visited[i] = True
                dfs(fatigue - dungeons[i][1], dungeon_count + 1)
                visited[i] = False
        return 
    
    dfs(k, 0)
    
    return max_dungeon_count
    
print(solution(80, [[80,20],[50,40],[30,10]]))
