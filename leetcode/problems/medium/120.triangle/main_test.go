package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMinimumTotal(t *testing.T) {
	assert.Equal(t, minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}), 11)
}
