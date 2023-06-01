package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
		{4, 5, 9},
		{5, 6, 11},
		{-1, 1, 0},
	}

	for _, test := range tests {
		if got := add(test.a, test.b); got != test.want {
			t.Errorf("add(%d, %d) = %d", test.a, test.b, got)
		}
	}

}

func add(a int, b int) int {
	return a + b
}
