package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGradingStudents(t *testing.T) {
	assert.ElementsMatch(t, gradingStudents([]int32{73, 67, 38, 33}), []int32{75, 67, 40, 33})
}
