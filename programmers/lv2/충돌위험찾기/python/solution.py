# https://school.programmers.co.kr/learn/courses/30/lessons/340211
# 프로그래머스의 난이도는 도저히 모르겠음... 쉬운문제가 없음...
class Robot:
    def __init__(self, path_coordinates):
        self._path_coordinates = path_coordinates
        self.current_row, self.current_col = self._path_coordinates[0]
        
        self._path_index = 1
        self._target_row, self._target_col = self._path_coordinates[self._path_index]
        
        self._is_finished = False

    def is_done(self):
        return self._is_finished

    def _update_target(self):
        self._path_index += 1
        
        if self._path_index < len(self._path_coordinates):
            self._target_row, self._target_col = self._path_coordinates[self._path_index]
        else:
            self._is_finished = True

    def move(self):
        if self.current_row != self._target_row:
            self.current_row += 1 if self.current_row < self._target_row else -1
        elif self.current_col != self._target_col:
            self.current_col += 1 if self.current_col < self._target_col else -1
        
        if self.current_row == self._target_row and self.current_col == self._target_col:
            self._update_target()

def collision_count(positions):
    if len(positions) == 0:
        return
    position_counts = {}
    for pos in positions:
        position_counts[pos] = position_counts.get(pos, 0) + 1
    collisions = 0

    for count in position_counts.values():
        if count > 1:
            collisions += 1
    return collisions

def solution(points, routes):
    robots = []
    max_total_move_seconds = 0

    for robot_route_points in routes:
        path_coords = [points[p_idx - 1] for p_idx in robot_route_points]
        robots.append(Robot(path_coords))
        
        total_move_seconds = 0
        for i in range(len(path_coords) - 1):
            start_posiotion = path_coords[i]
            target_position = path_coords[i+1]
            total_move_seconds += abs(start_posiotion[0] - target_position[0]) + abs(start_posiotion[1] - target_position[1])
        
        max_total_move_seconds = max(max_total_move_seconds, total_move_seconds)

    answer = 0
    
    initial_positions = [(r.current_row, r.current_col) for r in robots]
    answer += collision_count(initial_positions)

    for _ in range(max_total_move_seconds):
        positions_this_step = []
        for robot in robots:
            if robot.is_done():
                continue
            robot.move()
            
            positions_this_step.append((robot.current_row, robot.current_col))
        
        answer += collision_count(positions_this_step)
            
    return answer