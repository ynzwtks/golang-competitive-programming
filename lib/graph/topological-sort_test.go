package main

import (
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		name         string
		G            [][]int
		expectedSort []int
		expectedLen  int
	}{
		{
			name: "Test 1",
			G: [][]int{
				{},
				{2, 3},
				{3},
				{4},
				{},
			},
			expectedSort: []int{1, 2, 3, 4},
			expectedLen:  3,
		},
		{
			name: "Test 2",
			G: [][]int{
				{},
				{2},
				{3},
				{4},
				{5},
				{},
			},
			expectedSort: []int{1, 2, 3, 4, 5},
			expectedLen:  4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sort, longestPath := topologicalSort(tc.G)

			if !reflect.DeepEqual(sort, tc.expectedSort) {
				t.Errorf("Expected sort: %v, got: %v", tc.expectedSort, sort)
			}

			if longestPath != tc.expectedLen {
				t.Errorf("Expected longest path length: %d, got: %d", tc.expectedLen, longestPath)
			}
		})
	}
}
