package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.ElementsMatch(t, solution("2022.05.19", []string{"A 6", "B 12", "C 3"}, []string{"2021.05.02 A", "2021.07.01 B", "2022.02.19 C", "2022.02.20 C"}), []int{1, 3})
}
