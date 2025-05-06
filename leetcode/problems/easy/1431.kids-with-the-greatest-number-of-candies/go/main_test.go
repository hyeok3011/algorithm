package main

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	cases := []struct {
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
	}

	for _, c := range cases {
		result := kidsWithCandies(c.candies, c.extraCandies)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("kidsWithCandies(%v, %d) = %v; expected %v", c.candies, c.extraCandies, result, c.expected)
		}
	}
}
