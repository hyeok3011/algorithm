package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrayCode(t *testing.T) {
	assert.ElementsMatch(t, grayCode(2), []int{0, 1, 3, 2})
	assert.ElementsMatch(t, grayCode(1), []int{0, 1})
}
