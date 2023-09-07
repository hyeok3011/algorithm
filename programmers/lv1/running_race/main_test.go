package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	players := []string{"mumu", "soe", "poe", "kai", "mine"}
	calling := []string{"kai", "kai", "mine", "mine"}
	answer := []string{"mumu", "kai", "mine", "soe", "poe"}
	assert.ElementsMatch(t, solution(players, calling), answer)
}
