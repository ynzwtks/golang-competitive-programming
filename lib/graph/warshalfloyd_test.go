package main

import (
	"reflect"
	"testing"
)

func TestWarshallFloyd(t *testing.T) {
	n := 4

	edges := [][3]int{
		{0, 1, 2},
		{1, 2, 3},
		{1, 3, 5},
		{2, 3, 1},
		{3, 0, 10},
	}

	expected := [][]int{
		{0, 2, 5, 6},
		{14, 0, 3, 4},
		{11, 13, 0, 1},
		{10, 12, 15, 0},
	}

	wf := NewWarshallFloyd(n)

	for _, edge := range edges {
		wf.AddEdge(edge[0], edge[1], edge[2])
	}

	wf.Search()
	result := wf.GetResult()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
