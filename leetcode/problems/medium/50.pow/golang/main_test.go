package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMyPow(t *testing.T) {
	assert.Equal(t, myPow(2.0, 10), 1024.00000)
	// assert.Equal(t, myPow(2.1, 3), 9.26100)
	assert.Equal(t, myPow(2.0, -2), 0.25)
	assert.Equal(t, myPow(1, -2123124), 1.0)
	assert.Equal(t, myPow(-1, 2123124), -1.0)
}
