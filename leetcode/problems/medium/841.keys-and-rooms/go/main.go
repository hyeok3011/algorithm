package main

// https://leetcode.com/problems/keys-and-rooms/?envType=study-plan-v2&envId=leetcode-75
func canVisitAllRooms(rooms [][]int) bool {
    visited := make([]bool, len(rooms))

    var recursion func(roomIndex int)
    recursion = func(roomIndex int) {
        visited[roomIndex] = true
        for _, key := range rooms[roomIndex] {
            if visited[key] {
                continue
            }
            recursion(key)
        }
    }

    recursion(0)

    for i := range visited {
        if !visited[i] {
            return false
        }
    }
    return true
}
