package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t, climbStairs(2), 2)
	assert.Equal(t, climbStairs(3), 3)
	assert.Equal(t, climbStairs(4), 5)
	assert.Equal(t, climbStairs(5), 8)
	// 1111
	// 211ㄹ
	// 121
	// 112
	// 22

	// 5
	// 11111
	// 221
	// 122
	// 212
	// 2111
	// 1211
	// 1121
	// 1112

	// 6
	// 111111
	// 222
	// 21111
	// 12111
	// 11211
	// 11121
	// 11112
	// 2211
	// 2112
	// 1122
	// 1221

}
