package main

import (
	"reflect"
	"testing"
)

func TestEulerTour(t *testing.T) {
	g := [][]int{
		{},
		{2, 3},
		{4, 5},
		{6, 7},
		{},
		{},
		{},
		{},
	}

	eulerTour := NewEulerTour(g, 1)

	expectedInTime := []int{0, 0, 1, 7, 2, 4, 8, 10}
	expectedOutTime := []int{0, 13, 6, 12, 3, 5, 9, 11}
	expectedRoute := []int{1, 2, 4, 4, 5, 5, 2, 3, 6, 6, 7, 7, 3, 1}

	if !reflect.DeepEqual(eulerTour.inTime, expectedInTime) {
		t.Errorf("Expected inTime: %v, got: %v", expectedInTime, eulerTour.inTime)
	}

	if !reflect.DeepEqual(eulerTour.outTime, expectedOutTime) {
		t.Errorf("Expected outTime: %v, got: %v", expectedOutTime, eulerTour.outTime)
	}

	if !reflect.DeepEqual(eulerTour.route, expectedRoute) {
		t.Errorf("Expected route: %v, got: %v", expectedRoute, eulerTour.route)
	}

	expectedDepth := []int{0, 1, 2, 2, 1, 2, 2}
	expectedOrder := map[int]int{1: 0, 2: 1, 3: 4, 4: 2, 5: 3, 6: 5, 7: 6}

	order, depth := eulerTour.GetDepth()
	if !reflect.DeepEqual(order, expectedOrder) {
		t.Errorf("Expected order: %v, got: %v", expectedOrder, order)
	}

	if !reflect.DeepEqual(depth, expectedDepth) {
		t.Errorf("Expected depth: %v, got: %v", expectedDepth, depth)
	}

	for i := 1; i <= 7; i++ {
		inTime, outTime := eulerTour.GetRange(i)
		if inTime != expectedInTime[i] {
			t.Errorf("Node %d expected inTime: %d, got: %d", i, expectedInTime[i], inTime)
		}

		if outTime != expectedOutTime[i] {
			t.Errorf("Node %d expected outTime: %d, got: %d", i, expectedOutTime[i], outTime)
		}
	}
}
