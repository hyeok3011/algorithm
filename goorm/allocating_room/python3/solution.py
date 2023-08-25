# -*- coding: utf-8 -*-
# UTF-8 encoding when using korean
input_data = input()

splited_input_data = input_data.split(" ", -1)

room1, room2, room3 = int(splited_input_data[0]), int(splited_input_data[1]), int(splited_input_data[2])
person_count = int(splited_input_data[3])
memoization = set()

def allocate(person_count: int, room1: int, room2: int, room3: int, memoization: set) -> bool:
    if person_count == 0:
        return True

    if person_count < 0:
        return False

    if person_count in memoization:
        return False

    memoization.add(person_count)

    return allocate(person_count - room1,room1, room2, room3, memoization) or allocate(person_count - room2, room1, room2, room3, memoization) or allocate(person_count - room3, room1, room2, room3, memoization)

if allocate(person_count, room1, room2, room3, memoization):
    print(1)
else:
    print(0)
