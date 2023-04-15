package main

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	g := [][]int{
		{1, 2},
		{0, 2},
		{0, 1},
		{4},
		{3},
	}
	visited, prev := bfs(g, 0)

	expectedVisited := []int{0, 1, 1, -1, -1}
	expectedPath := []int{0, 1}

	if !reflect.DeepEqual(visited, expectedVisited) {
		t.Errorf("Expected visited %v, but got %v", expectedVisited, visited)
	}

	path := restoreBfsPath(prev, 0, 1)

	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Expected path %v, but got %v", expectedPath, path)
	}
}

func TestRestorePath(t *testing.T) {
	prev := []int{-1, 0, 0, -1, -1}
	path := restoreBfsPath(prev, 0, 1)

	expectedPath := []int{0, 1}

	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Expected path %v, but got %v", expectedPath, path)
	}
}

func TestCreateBFSGroup(t *testing.T) {
	g := [][]int{
		{1, 2},
		{0, 2},
		{0, 1},
		{4},
		{3},
	}
	groups := createBfsGroup(g, 4) // 修正: 引数 n の値を 5 から 4 に変更

	expectedGroups := map[int]int{0: 3, 3: 2}

	if !reflect.DeepEqual(groups, expectedGroups) {
		t.Errorf("Expected groups %v, but got %v", expectedGroups, groups)
	}
}
