package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotationLeft(t *testing.T) {
	assert.ElementsMatch(t, rotateLeft(4, []int32{1, 2, 3, 4, 5}), []int32{5, 1, 2, 3, 4})
	assert.ElementsMatch(t, rotateLeft(2, []int32{1, 2, 3, 4, 5}), []int32{3, 4, 5, 1, 2})
}
