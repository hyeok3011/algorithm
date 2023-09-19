package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoluiton(t *testing.T) {

	assert.ElementsMatch(t, solution([]string{"may", "kein", "kain", "radi"}, []int{5, 10, 1, 3}, [][]string{
		{"may", "kein", "kain", "radi"},
		{"may", "kein", "brin", "deny"},
		{"kon", "kain", "may", "coni"},
	}), []int{19, 15, 6})
}
