package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestSolution(t *testing.T) {
	// assert.Equal(t, solution(28, 18, 26, 10, 8,
	// 	[]int{0, 0, 1, 1, 1, 1, 1}), 40)
	// assert.Equal(t, solution(-10, -5, 5, 5, 1, []int{0, 0, 0, 0, 0, 1, 0}), 25)
	// assert.Equal(t, solution(11, 8, 10, 10, 1, []int{0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1}), 20)
	assert.Equal(t, solution(11, 8, 10, 10, 100, []int{0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1}), 60)
}

// target = max(t1 - temperature, temperature - t2)
// 18 26
// 0	x	28	24
// 1	x	27	24
// 2	o	26	24 26
// 3	o	25	24
// 4	o	24	off
// 5	o	25	off
// 6	o	26	off

// -5 5
// 0	x	-10	40
// 1	x	-9	40
// 2	x	-8	40
// 3	x	-7	40
// 4	x	-6	40
// 5	o	-5	off
// 6	x	-6	off

// 8 10
// 0	x	11	10
// 1	o	10	10
// 2	o	10	10
// 3	o	10	10
// 4	o	10	10
// 5	o	10	10
// 6	o	10	10
// 7	x	10	10
// 8	x	10	10
// 9	x	10	10
// 10	o	10	10
// 11	o	10	off
