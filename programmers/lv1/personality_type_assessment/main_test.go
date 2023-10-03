package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]string{"AN", "CF", "MJ", "RT", "NA"}, []int{5, 3, 2, 7, 5}), "TCMA")
	assert.Equal(t, solution([]string{"TR", "RT", "TR"}, []int{7, 1, 3}), "RCJA")
}
