package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	nums := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums, 3, nums2, 3)
	assert.ElementsMatch(t, nums, []int{1, 2, 2, 3, 5, 6})
}
