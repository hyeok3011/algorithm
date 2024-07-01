package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)
	// 4,-1,2,1
	assert.Equal(t, maxSubArray([]int{5, 4, -1, 7, 8}), 23)
	// 5,4,-1,7,8
	assert.Equal(t, maxSubArray([]int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}), 6)
	// 2, 1, -2, 1, 4

}
