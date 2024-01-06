package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, solution([]string{"muzi", "ryan", "frodo", "neo"}, []string{"muzi frodo", "muzi frodo", "ryan muzi", "ryan muzi", "ryan muzi", "frodo muzi", "frodo ryan", "neo muzi"}), 2)
	// assert.Equal(t, solution([]string{"joy", "brad", "alessandro", "conan", "david"}, []string{"alessandro brad", "alessandro joy", "alessandro conan", "david alessandro", "alessandro david"}), 4)
	// assert.Equal(t, solution([]string{"a", "b", "c"}, []string{"a b", "b a", "c a", "a c", "a c", "c a"}), 0)

}
