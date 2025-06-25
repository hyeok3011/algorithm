# https://school.programmers.co.kr/learn/courses/30/lessons/160586
def solution(keymap, targets):
    key_cost_map = {}
    for key_layout in keymap:
        for index, char in enumerate(key_layout):
            press_count = index + 1
            key_cost_map[char] = min(key_cost_map.get(char, 101), press_count)

    results = []
    for target_str in targets:
        total_presses = 0
        is_possible = True

        for char in target_str:
            if char not in key_cost_map:
                is_possible = False
                break

            total_presses += key_cost_map[char]

        if is_possible:
            results.append(total_presses)
        else:
            results.append(-1)

    return results