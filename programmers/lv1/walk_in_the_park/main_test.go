package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	park := []string{"SOO", "OOO", "OOO"}
	routes := []string{"E 2", "S 2", "W 1"}
	assert.ElementsMatch(t, solution(park, routes), []int{2, 1})

	park2 := []string{"SOO", "OXX", "OOO"}
	routes2 := []string{"E 2", "S 2", "W 1"}
	assert.ElementsMatch(t, solution(park2, routes2), []int{0, 1})
}
