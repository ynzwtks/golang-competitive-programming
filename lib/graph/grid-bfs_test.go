package main

import (
	"reflect"
	"testing"
)

func TestGridBfs(t *testing.T) {
	grid := [][]byte{
		{'S', '.', '.', '#', 'G'},
		{'#', '#', '.', '.', '#'},
		{'.', '.', '.', '.', '.'},
		{'.', '#', '#', '.', '#'},
		{'.', '.', '.', '.', '.'},
	}

	expected := [][]int{
		{0, 1, 2, -1, -1},
		{-1, -1, 3, 4, -1},
		{6, 5, 4, 5, 6},
		{7, -1, -1, 6, -1},
		{8, 9, 8, 7, 8},
	}

	result := gridBfs(grid, 0, 0)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
