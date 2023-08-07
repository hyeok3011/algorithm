package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduc(t *testing.T) {
	// assert.Equal(t, maxProduct2([]int{2, 3, -2, 4}), 6)
	// assert.Equal(t, maxProduct2([]int{-2, 0, -1}), 0)
	// assert.Equal(t, maxProduct2([]int{-2, 3, -4}), 24)
	assert.Equal(t, maxProduct2([]int{3, -1, 4}), 4)
	assert.Equal(t, maxProduct2([]int{2, -5, -2, -4, 3}), 24)
	assert.Equal(t, maxProduct2([]int{-1, -2, -3, 0}), 6)

}
