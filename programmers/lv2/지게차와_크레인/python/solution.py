# https://school.programmers.co.kr/learn/courses/30/lessons/388353?language=python3
EMPTY_SLOT = '0'

def solution(storage, requests):
    storage = [list(row) for row in storage]
    rows, cols = len(storage), len(storage[0])

    container_map = {}
    for r in range(rows):
        for c in range(cols):
            container_id = storage[r][c]
            if container_id == EMPTY_SLOT:
                continue
            container_map.setdefault(container_id, set()).add((r, c))

    for req in requests:
        target_id = req[0]
        is_crane_request = len(req) > 1

        if target_id not in container_map or not container_map[target_id]:
            continue

        positions_to_remove = set()

        if is_crane_request:
            positions_to_remove = container_map[target_id].copy()
        else:
            containers_to_check = list(container_map[target_id])

            for r, c in containers_to_check:
                path_cache = {}
                original_value = storage[r][c]
                storage[r][c] = EMPTY_SLOT

                if can_reach_outside(r, c, storage, path_cache, set()):
                    positions_to_remove.add((r, c))

                storage[r][c] = original_value

        if positions_to_remove:
            for r, c in positions_to_remove:
                storage[r][c] = EMPTY_SLOT
                container_map[target_id].remove((r, c))

    remaining_count = 0
    for r in range(rows):
        for c in range(cols):
            if storage[r][c] != EMPTY_SLOT:
                remaining_count += 1

    return remaining_count

def can_reach_outside(r, c, storage, cache, visited):
    pos = (r, c)
    rows, cols = len(storage), len(storage[0])

    if not (0 <= r < rows and 0 <= c < cols):
        return True

    if pos in cache:
        return cache[pos]

    if pos in visited or storage[r][c] != EMPTY_SLOT:
        return False

    visited.add(pos)

    is_path_found = (can_reach_outside(r - 1, c, storage, cache, visited) or
                     can_reach_outside(r + 1, c, storage, cache, visited) or
                     can_reach_outside(r, c - 1, storage, cache, visited) or
                     can_reach_outside(r, c + 1, storage, cache, visited))

    visited.remove(pos)

    cache[pos] = is_path_found
    return is_path_found