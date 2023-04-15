package main

import (
	"math"
	"testing"
)

func TestNewBellmanFord(t *testing.T) {
	bf := NewBellmanFord(5)
	if bf.n != 5 {
		t.Errorf("Expected n to be 5, got %d", bf.n)
	}
}

func TestAddEdge(t *testing.T) {
	bf := NewBellmanFord(3)
	bf.AddEdge(0, 1, 2)
	bf.AddEdge(1, 2, 3)

	if len(bf.edges) != 2 {
		t.Errorf("Expected number of edges to be 2, got %d", len(bf.edges))
	}

	expectedEdge := Edge{0, 1, 2}
	if bf.edges[0] != expectedEdge {
		t.Errorf("Expected edge to be %v, got %v", expectedEdge, bf.edges[0])
	}
}

func TestSearch(t *testing.T) {
	bf := NewBellmanFord(4)
	bf.AddEdge(0, 1, 1)
	bf.AddEdge(0, 2, 5)
	bf.AddEdge(1, 2, 2)
	bf.AddEdge(1, 3, 5)
	bf.AddEdge(2, 3, 1)

	dist := bf.Search(0)
	expectedDist := []int{0, 1, 3, 4}
	for i, d := range dist {
		if d != expectedDist[i] {
			t.Errorf("Expected dist[%d] to be %d, got %d", i, expectedDist[i], d)
		}
	}

	bf = NewBellmanFord(3)
	bf.AddEdge(0, 1, 1)
	bf.AddEdge(1, 2, 1)
	bf.AddEdge(2, 0, -3)

	dist = bf.Search(0)
	if dist != nil {
		t.Errorf("Expected dist to be nil due to negative cycle, got %v", dist)
	}
}

func TestSearchWithNegativeWeights(t *testing.T) {
	bf := NewBellmanFord(4)
	bf.AddEdge(0, 1, 2)
	bf.AddEdge(0, 2, 5)
	bf.AddEdge(1, 2, -4)
	bf.AddEdge(1, 3, 5)
	bf.AddEdge(2, 3, 1)

	dist := bf.Search(0)
	expectedDist := []int{0, 2, -2, -1}
	for i, d := range dist {
		if d != expectedDist[i] {
			t.Errorf("Expected dist[%d] to be %d, got %d", i, expectedDist[i], d)
		}
	}
}

func TestSearchWithUnreachableNodes(t *testing.T) {
	bf := NewBellmanFord(5)
	bf.AddEdge(0, 1, 1)
	bf.AddEdge(1, 2, 2)
	bf.AddEdge(3, 4, 3)

	dist := bf.Search(0)
	expectedDist := []int{0, 1, 3, math.MaxInt64, math.MaxInt64}
	for i, d := range dist {
		if d != expectedDist[i] {
			t.Errorf("Expected dist[%d] to be %d, got %d", i, expectedDist[i], d)
		}
	}
}
