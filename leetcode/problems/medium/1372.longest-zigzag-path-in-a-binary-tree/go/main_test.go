package main

import (
	"testing"

	"github.com/go-playground/assert"
	. "github.com/hyeok3011/algorithm/structures"
)

func TestLongestZigZag(t *testing.T) {
	assert.Equal(t, longestZigZag(BuildTreeFromArray([]any{1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1})), 3)
	assert.Equal(t, longestZigZagBottomUp(BuildTreeFromArray([]any{1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1})), 3)
}
