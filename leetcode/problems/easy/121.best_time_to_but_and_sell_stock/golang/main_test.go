package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
	assert.Equal(t, maxProfit([]int{7, 6, 4, 3, 1}), 0)
	assert.Equal(t, maxProfit([]int{7, 6, 4, 3, 1}), 0)
	assert.Equal(t, maxProfit([]int{3, 2, 6, 5, 0, 3}), 4)
}
