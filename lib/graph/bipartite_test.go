package main

import (
	"testing"
)

func TestNewBipartite(t *testing.T) {
	n := 4
	l := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}

	b := NewBipartite(n, l)

	if b.n != n {
		t.Errorf("Expected %d, but got %d", n, b.n)
	}

	if !b.IsBipartite() {
		t.Errorf("Expected the graph to be bipartite")
	}
}

func TestIsBipartite(t *testing.T) {
	testCases := []struct {
		n        int
		l        [][]int
		expected bool
	}{
		{
			n: 4,
			l: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
			},
			expected: true,
		},
		{
			n: 3,
			l: [][]int{
				{0, 1},
				{1, 2},
				{2, 0},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		b := NewBipartite(tc.n, tc.l)
		if result := b.IsBipartite(); result != tc.expected {
			t.Errorf("Expected %t, but got %t", tc.expected, result)
		}
	}
}
