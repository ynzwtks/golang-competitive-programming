package main

import (
	"testing"
)

func TestMST(t *testing.T) {
	mst := NewMST(6)
	edges := []Edge{
		{1, 2, 1},
		{1, 3, 2},
		{2, 4, 3},
		{2, 5, 4},
		{3, 5, 5},
	}

	for _, edge := range edges {
		mst.AddEdge(edge.source, edge.destination, edge.weight)
	}

	cost := mst.Build(true)
	expectedCost := 10
	if cost != expectedCost {
		t.Errorf("Expected cost: %d, got: %d", expectedCost, cost)
	}

}
