package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution([]string{".#...", "..#..", "...#."}), []int{0, 1, 3, 4})
	assert.ElementsMatch(t, solution([]string{"..........", ".....#....", "......##..", "...##.....", "....#....."}), []int{1, 3, 5, 8})
	assert.ElementsMatch(t, solution([]string{".##...##.", "#..#.#..#", "#...#...#", ".#.....#.", "..#...#..", "...#.#...", "....#...."}), []int{0, 0, 7, 9})
	assert.ElementsMatch(t, solution([]string{"..", "#."}), []int{1, 0, 2, 1})
}
