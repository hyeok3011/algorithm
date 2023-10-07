package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolutuion(t *testing.T) {
	assert.Equal(t, solution([]int{180, 5000, 10, 600}, []string{"05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT", "07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT"}), []int{14600, 34400, 5000})
	assert.Equal(t, solution([]int{120, 0, 60, 591}, []string{"16:00 3961 IN", "16:00 0202 IN", "18:00 3961 OUT", "18:00 0202 OUT", "23:58 3961 IN"}), []int{0, 591})
}
