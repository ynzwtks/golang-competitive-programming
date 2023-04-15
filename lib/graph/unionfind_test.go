package main

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(5)

	if uf.group != 5 {
		t.Errorf("Expected 5 groups, but got %d", uf.group)
	}

	uf.Union(0, 1)

	if uf.group != 4 {
		t.Errorf("Expected 4 groups, but got %d", uf.group)
	}

	if !uf.Connected(0, 1) {
		t.Error("Expected 0 and 1 to be connected")
	}

	uf.Union(1, 2)

	if uf.group != 3 {
		t.Errorf("Expected 3 groups, but got %d", uf.group)
	}

	if !uf.Connected(0, 2) {
		t.Error("Expected 0 and 2 to be connected")
	}

	if uf.Connected(0, 3) {
		t.Error("Expected 0 and 3 to not be connected")
	}

	groups := uf.Groups()
	expectedGroups := map[int]int{
		uf.find(0): 3,
		uf.find(3): 1,
		uf.find(4): 1,
	}

	for k, v := range expectedGroups {
		if groups[k] != v {
			t.Errorf("Expected group %d to have %d elements, but got %d", k, v, groups[k])
		}
	}
}
