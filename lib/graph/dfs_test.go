package main

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
	G := [][]int{
		{1, 2},
		{2, 3},
		{3},
		{},
	}

	dfs := NewDFS(G)

	found := dfs.Search(0, 3)

	if !found {
		t.Errorf("Expected to find a path from 0 to 3, but no path found")
	}

	expectedShortestPath := []int{0, 1, 3}
	if !reflect.DeepEqual(dfs.shortestPath, expectedShortestPath) {
		t.Errorf("Expected shortest path %v but got %v", expectedShortestPath, dfs.shortestPath)
	}
}
