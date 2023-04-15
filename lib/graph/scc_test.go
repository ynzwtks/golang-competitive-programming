package main

import (
	"reflect"
	"testing"
)

func TestNewScc(t *testing.T) {
	g := [][]int{
		{1},
		{2},
		{0},
		{4},
		{3},
	}
	scc := NewScc(g)

	if scc.nodes != 5 {
		t.Errorf("Expected scc.nodes to be 5, but got %d", scc.nodes)
	}

	if !reflect.DeepEqual(scc.edges, g) {
		t.Errorf("Expected scc.edges to be equal to g, but got %v", scc.edges)
	}
}

func TestScc(t *testing.T) {
	g := [][]int{
		{1},
		{2},
		{0},
		{4},
		{3},
	}
	scc := NewScc(g)
	sccs := scc.scc()

	expectedSCCs := [][]int{
		{0, 2, 1},
		{3, 4},
	}

	if len(sccs) != len(expectedSCCs) {
		t.Fatalf("Expected %d strongly connected components, but got %d", len(expectedSCCs), len(sccs))
	}

	for i, scc := range sccs {
		if len(scc) != len(expectedSCCs[i]) {
			t.Errorf("Expected SCC %d to have length %d, but got %d", i, len(expectedSCCs[i]), len(scc))
		}

		for _, v := range expectedSCCs[i] {
			found := false
			for _, w := range scc {
				if v == w {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected vertex %d to be in SCC %d, but it was not", v, i)
			}
		}
	}
}
