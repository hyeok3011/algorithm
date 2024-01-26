package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingLeaderBoard(t *testing.T) {
	assert.ElementsMatch(t, climbingLeaderboard([]int32{100, 90, 90, 80}, []int32{70, 80, 80, 105}), []int32{4, 3, 3, 1})
}
