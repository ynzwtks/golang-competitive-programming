package main

import (
	"fmt"
	"testing"
)

func TestWeightedUnionFind(t *testing.T) {
	wuf := NewWeightedUnionFind(10)

	// Test Root function
	if root := wuf.Root(0); root != 0 {
		t.Errorf("Root of 0 should be 0, got %d", root)
	}

	// Test Unite function
	united := wuf.Unite(1, 5, 2)
	if !united {
		t.Errorf("Unite(1, 5, 2) should return true, got false")
	}

	// Test Same function
	if !wuf.Same(1, 5) {
		t.Errorf("1 and 5 should be in the same group")
	}

	// Test Weight function
	if weight := wuf.Diff(1, 5); weight != 2 {
		t.Errorf("Weight of 5 should be 2, got %d", weight)
	}

	// Test Unite function again
	united = wuf.Unite(5, 6, 2)
	if !united {
		t.Errorf("Unite(5, 6, 2) should return true, got false")
	}

	// Test Diff function
	if diff := wuf.Diff(1, 6); diff != 4 {
		t.Errorf("Diff(1, 6) should be 4, got %d", diff)
	}

	// Test Same function again
	if !wuf.Same(1, 6) {
		t.Errorf("1 and 6 should be in the same group")
	}

	// Test a more complex Unite pattern
	wuf.Unite(2, 7, 3)
	wuf.Unite(7, 8, 1)
	wuf.Unite(1, 2, 5)

	if !wuf.Same(2, 8) {
		t.Errorf("2 and 8 should be in the same group")
	}

	if diff := wuf.Diff(2, 8); diff != 4 {
		t.Errorf("Diff(2, 8) should be 4, got %d", diff)
	}

	if diff := wuf.Diff(8, 1); diff != -9 {
		t.Errorf("Diff(1, 2) should be -9, got %d", diff)
	}

	if !wuf.Same(1, 7) {
		t.Errorf("1 and 7 should be in the same group")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Running WeightedUnionFind tests")
	exitVal := m.Run()
	if exitVal == 0 {
		fmt.Println("All tests passed")
	} else {
		fmt.Println("Some tests failed")
	}
}
